package test

import (
	"context"
	"fmt"

	"github.com/spiffe/spire/pkg/common/catalog/interfaces"
)

func NewTestHostService() TestHostService {
	return &testHostService{}
}

type testHostService struct{}

func (*testHostService) CallHostService(ctx context.Context, req *Request) (*Response, error) {
	pluginName, ok := interfaces.PluginNameFromHostServiceContext(ctx)
	if !ok {
		pluginName = "<unknown>"
	}
	return &Response{
		Out: fmt.Sprintf("hostservice[plugin=%s](%s)", pluginName, req.In),
	}, nil
}
