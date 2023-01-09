// Package serverdiscovery provides a service on Consul servers to discover the set of servers
// currently able to handle incoming requests.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: proto-public/pbserverdiscovery/serverdiscovery.proto

package pbserverdiscovery

import (
	_ "github.com/hashicorp/consul/proto-public/annotations/ratelimit"
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

type WatchServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Wan being set to true will cause WAN addresses to be sent in the response
	// instead of the LAN addresses which are the default
	Wan bool `protobuf:"varint,1,opt,name=wan,proto3" json:"wan,omitempty"`
}

func (x *WatchServersRequest) Reset() {
	*x = WatchServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchServersRequest) ProtoMessage() {}

func (x *WatchServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchServersRequest.ProtoReflect.Descriptor instead.
func (*WatchServersRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescGZIP(), []int{0}
}

func (x *WatchServersRequest) GetWan() bool {
	if x != nil {
		return x.Wan
	}
	return false
}

type WatchServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Servers is the list of server address information.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *WatchServersResponse) Reset() {
	*x = WatchServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchServersResponse) ProtoMessage() {}

func (x *WatchServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchServersResponse.ProtoReflect.Descriptor instead.
func (*WatchServersResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescGZIP(), []int{1}
}

func (x *WatchServersResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique string identifying this server for all time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// address on the network of the server
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// the consul version of the server
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Server) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Server) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proto_public_pbserverdiscovery_serverdiscovery_proto protoreflect.FileDescriptor

var file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x13,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x77, 0x61, 0x6e, 0x22, 0x5a, 0x0a, 0x14, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x22, 0x4c, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32,
	0xa2, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x35, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0xe2, 0x86, 0x04, 0x02,
	0x08, 0x02, 0x30, 0x01, 0x42, 0x9a, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x42, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x53, 0xaa, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xca, 0x02, 0x20, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xe2, 0x02, 0x2c,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescOnce sync.Once
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescData = file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDesc
)

func file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescGZIP() []byte {
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescOnce.Do(func() {
		file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescData)
	})
	return file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDescData
}

var file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_public_pbserverdiscovery_serverdiscovery_proto_goTypes = []interface{}{
	(*WatchServersRequest)(nil),  // 0: hashicorp.consul.serverdiscovery.WatchServersRequest
	(*WatchServersResponse)(nil), // 1: hashicorp.consul.serverdiscovery.WatchServersResponse
	(*Server)(nil),               // 2: hashicorp.consul.serverdiscovery.Server
}
var file_proto_public_pbserverdiscovery_serverdiscovery_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.serverdiscovery.WatchServersResponse.servers:type_name -> hashicorp.consul.serverdiscovery.Server
	0, // 1: hashicorp.consul.serverdiscovery.ServerDiscoveryService.WatchServers:input_type -> hashicorp.consul.serverdiscovery.WatchServersRequest
	1, // 2: hashicorp.consul.serverdiscovery.ServerDiscoveryService.WatchServers:output_type -> hashicorp.consul.serverdiscovery.WatchServersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_public_pbserverdiscovery_serverdiscovery_proto_init() }
func file_proto_public_pbserverdiscovery_serverdiscovery_proto_init() {
	if File_proto_public_pbserverdiscovery_serverdiscovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchServersRequest); i {
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
		file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchServersResponse); i {
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
		file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
			RawDescriptor: file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_public_pbserverdiscovery_serverdiscovery_proto_goTypes,
		DependencyIndexes: file_proto_public_pbserverdiscovery_serverdiscovery_proto_depIdxs,
		MessageInfos:      file_proto_public_pbserverdiscovery_serverdiscovery_proto_msgTypes,
	}.Build()
	File_proto_public_pbserverdiscovery_serverdiscovery_proto = out.File
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_rawDesc = nil
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_goTypes = nil
	file_proto_public_pbserverdiscovery_serverdiscovery_proto_depIdxs = nil
}
