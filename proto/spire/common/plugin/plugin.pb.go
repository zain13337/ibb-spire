// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/common/plugin/plugin.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//* Represents the plugin-specific configuration string.
type ConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* The configuration for the plugin.
	Configuration string `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	//* Global configurations.
	GlobalConfig *ConfigureRequest_GlobalConfig `protobuf:"bytes,2,opt,name=globalConfig,proto3" json:"globalConfig,omitempty"`
}

func (x *ConfigureRequest) Reset() {
	*x = ConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRequest) ProtoMessage() {}

func (x *ConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRequest.ProtoReflect.Descriptor instead.
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigureRequest) GetConfiguration() string {
	if x != nil {
		return x.Configuration
	}
	return ""
}

func (x *ConfigureRequest) GetGlobalConfig() *ConfigureRequest_GlobalConfig {
	if x != nil {
		return x.GlobalConfig
	}
	return nil
}

//* Represents a list of configuration problems
//found in the configuration string.
type ConfigureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* A list of errors
	ErrorList []string `protobuf:"bytes,1,rep,name=errorList,proto3" json:"errorList,omitempty"`
}

func (x *ConfigureResponse) Reset() {
	*x = ConfigureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureResponse) ProtoMessage() {}

func (x *ConfigureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureResponse.ProtoReflect.Descriptor instead.
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigureResponse) GetErrorList() []string {
	if x != nil {
		return x.ErrorList
	}
	return nil
}

//* Represents an empty request.
type GetPluginInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginInfoRequest) Reset() {
	*x = GetPluginInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginInfoRequest) ProtoMessage() {}

func (x *GetPluginInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{2}
}

//* Represents the plugin metadata.
type GetPluginInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DateCreated string `protobuf:"bytes,5,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	Location    string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Version     string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Author      string `protobuf:"bytes,8,opt,name=author,proto3" json:"author,omitempty"`
	Company     string `protobuf:"bytes,9,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *GetPluginInfoResponse) Reset() {
	*x = GetPluginInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginInfoResponse) ProtoMessage() {}

func (x *GetPluginInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginInfoResponse.ProtoReflect.Descriptor instead.
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *GetPluginInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPluginInfoResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetPluginInfoResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetPluginInfoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetPluginInfoResponse) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *GetPluginInfoResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetPluginInfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetPluginInfoResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetPluginInfoResponse) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostServices []string `protobuf:"bytes,1,rep,name=host_services,json=hostServices,proto3" json:"host_services,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *InitRequest) GetHostServices() []string {
	if x != nil {
		return x.HostServices
	}
	return nil
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginServices []string `protobuf:"bytes,1,rep,name=plugin_services,json=pluginServices,proto3" json:"plugin_services,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *InitResponse) GetPluginServices() []string {
	if x != nil {
		return x.PluginServices
	}
	return nil
}

//* Global configuration nested type.
type ConfigureRequest_GlobalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrustDomain string `protobuf:"bytes,1,opt,name=trustDomain,proto3" json:"trustDomain,omitempty"`
}

func (x *ConfigureRequest_GlobalConfig) Reset() {
	*x = ConfigureRequest_GlobalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_common_plugin_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureRequest_GlobalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRequest_GlobalConfig) ProtoMessage() {}

func (x *ConfigureRequest_GlobalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_spire_common_plugin_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRequest_GlobalConfig.ProtoReflect.Descriptor instead.
func (*ConfigureRequest_GlobalConfig) Descriptor() ([]byte, []int) {
	return file_spire_common_plugin_plugin_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigureRequest_GlobalConfig) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

var File_spire_common_plugin_plugin_proto protoreflect.FileDescriptor

var file_spire_common_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x30, 0x0a, 0x0c, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x31, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x87, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x22, 0x32, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0x59,
	0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x4b, 0x0a, 0x04,
	0x49, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_common_plugin_plugin_proto_rawDescOnce sync.Once
	file_spire_common_plugin_plugin_proto_rawDescData = file_spire_common_plugin_plugin_proto_rawDesc
)

func file_spire_common_plugin_plugin_proto_rawDescGZIP() []byte {
	file_spire_common_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_spire_common_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_common_plugin_plugin_proto_rawDescData)
	})
	return file_spire_common_plugin_plugin_proto_rawDescData
}

var file_spire_common_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spire_common_plugin_plugin_proto_goTypes = []interface{}{
	(*ConfigureRequest)(nil),              // 0: spire.common.plugin.ConfigureRequest
	(*ConfigureResponse)(nil),             // 1: spire.common.plugin.ConfigureResponse
	(*GetPluginInfoRequest)(nil),          // 2: spire.common.plugin.GetPluginInfoRequest
	(*GetPluginInfoResponse)(nil),         // 3: spire.common.plugin.GetPluginInfoResponse
	(*InitRequest)(nil),                   // 4: spire.common.plugin.InitRequest
	(*InitResponse)(nil),                  // 5: spire.common.plugin.InitResponse
	(*ConfigureRequest_GlobalConfig)(nil), // 6: spire.common.plugin.ConfigureRequest.GlobalConfig
}
var file_spire_common_plugin_plugin_proto_depIdxs = []int32{
	6, // 0: spire.common.plugin.ConfigureRequest.globalConfig:type_name -> spire.common.plugin.ConfigureRequest.GlobalConfig
	4, // 1: spire.common.plugin.PluginInit.Init:input_type -> spire.common.plugin.InitRequest
	5, // 2: spire.common.plugin.PluginInit.Init:output_type -> spire.common.plugin.InitResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_common_plugin_plugin_proto_init() }
func file_spire_common_plugin_plugin_proto_init() {
	if File_spire_common_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_common_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_common_plugin_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureRequest_GlobalConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_common_plugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_common_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_spire_common_plugin_plugin_proto_depIdxs,
		MessageInfos:      file_spire_common_plugin_plugin_proto_msgTypes,
	}.Build()
	File_spire_common_plugin_plugin_proto = out.File
	file_spire_common_plugin_plugin_proto_rawDesc = nil
	file_spire_common_plugin_plugin_proto_goTypes = nil
	file_spire_common_plugin_plugin_proto_depIdxs = nil
}
