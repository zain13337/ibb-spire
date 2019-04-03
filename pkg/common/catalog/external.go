package catalog

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/log"
	"github.com/zeebo/errs"
	"google.golang.org/grpc"
)

const (
	// the ID used to dial host services
	hostServicesID = 1
)

func PluginMain(plugin Plugin) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: goplugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   plugin.Plugin.PluginType(),
			MagicCookieValue: plugin.Plugin.PluginType(),
		},
		Plugins: map[string]goplugin.Plugin{
			plugin.Name: &hcServerPlugin{
				logger: logger,
				plugin: plugin,
			},
		},
		Logger:     logger,
		GRPCServer: goplugin.DefaultGRPCServer,
	})
}

type hcServerPlugin struct {
	goplugin.NetRPCUnsupportedPlugin
	logger hclog.Logger
	plugin Plugin
}

var _ goplugin.GRPCPlugin = (*hcServerPlugin)(nil)

func (p *hcServerPlugin) GRPCServer(b *goplugin.GRPCBroker, s *grpc.Server) error {
	initPluginServer(
		s,
		grpcBrokerDialer{b: b},
		p.logger,
		p.plugin.Plugin,
		p.plugin.Services,
	)
	return nil
}

func (p *hcServerPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("unimplemented")
}

type grpcBrokerDialer struct {
	b *goplugin.GRPCBroker
}

func (d grpcBrokerDialer) DialHost() (*grpc.ClientConn, error) {
	return d.b.Dial(hostServicesID)
}

type ExternalPlugin struct {
	Log           logrus.FieldLogger
	Name          string
	Path          string
	Checksum      string
	Data          string
	Plugin        PluginClient
	KnownServices []ServiceClient
	HostServices  []HostServiceServer
}

func LoadExternalPlugin(ctx context.Context, ext ExternalPlugin) (plugin *CatalogPlugin, err error) {
	// Resolve to an absolute path. We don't to use path environment lookups.
	ext.Path, err = filepath.Abs(ext.Path)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	cmd := pluginCmd(ext.Path)

	var secureConfig *goplugin.SecureConfig
	if ext.Checksum != "" {
		secureConfig, err = buildSecureConfig(ext.Checksum)
		if err != nil {
			return nil, err
		}
	} else {
		ext.Log.Warn("Plugin not using secure ext")
	}

	logger := log.HCLogAdapter{
		Log:  ext.Log,
		Name: "external",
	}

	// start the external plugin. ensure it is killed if there is an error.
	pluginClient := goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: goplugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   ext.Plugin.PluginType(),
			MagicCookieValue: ext.Plugin.PluginType(),
		},
		Cmd: cmd,
		// TODO: enable AutoMTLS if it is fixed to work with brokering.
		//AutoMTLS:         true,
		AllowedProtocols: []goplugin.Protocol{goplugin.ProtocolGRPC},
		Plugins: map[string]goplugin.Plugin{
			"external": &hcClientPlugin{
				ext: ext,
			},
		},
		Logger:       logger.Named(ext.Name),
		SecureConfig: secureConfig,
	})
	defer func() {
		if err != nil {
			pluginClient.Kill()
		}
	}()

	// create the GRPC client and ensure it is closed on error
	grpcClient, err := pluginClient.Client()
	if err != nil {
		return nil, err
	}

	// the primary interface is dispensed via the plugin name
	pluginRaw, err := grpcClient.Dispense("external")
	if err != nil {
		return nil, err
	}

	plugin, ok := pluginRaw.(*CatalogPlugin)
	if !ok {
		// shouldn't happen.
		return nil, errs.New("expected %T, got %T", plugin, pluginRaw)
	}

	// Kill also closes the gRPC client
	plugin.closer = pluginClient.Kill

	return plugin, nil
}

func buildSecureConfig(checksum string) (*goplugin.SecureConfig, error) {
	sum, err := hex.DecodeString(checksum)
	if err != nil {
		return nil, fmt.Errorf("unable to decode checksum: %v", err)
	}

	return &goplugin.SecureConfig{
		Checksum: sum,
		Hash:     sha256.New(),
	}, nil
}

type hcClientPlugin struct {
	goplugin.NetRPCUnsupportedPlugin
	ext ExternalPlugin
}

var _ goplugin.GRPCPlugin = (*hcClientPlugin)(nil)

func (p *hcClientPlugin) GRPCServer(b *goplugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("not implemented host side")
}

func (p *hcClientPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	// kick off a goroutine that serves the host services to the plugin
	go b.AcceptAndServe(hostServicesID, func(opts []grpc.ServerOption) *grpc.Server {
		return NewHostServer(p.ext.Name, opts, p.ext.HostServices)
	})

	plugin, err := newCatalogPlugin(ctx, c, catalogPluginConfig{
		Log:           p.ext.Log,
		Name:          p.ext.Name,
		Plugin:        p.ext.Plugin,
		KnownServices: p.ext.KnownServices,
		HostServices:  p.ext.HostServices,
	})
	if err != nil {
		return nil, err
	}
	return plugin, nil
}
