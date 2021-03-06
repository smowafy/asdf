// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: vault.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Vault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents map[string][]byte `protobuf:"bytes,1,rep,name=Contents,proto3" json:"Contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Vault) Reset() {
	*x = Vault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vault) ProtoMessage() {}

func (x *Vault) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vault.ProtoReflect.Descriptor instead.
func (*Vault) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{0}
}

func (x *Vault) GetContents() map[string][]byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_vault_proto protoreflect.FileDescriptor

var file_vault_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x73, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7c, 0x0a, 0x05, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x89, 0x01,
	0x0a, 0x0b, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3c, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_proto_rawDescOnce sync.Once
	file_vault_proto_rawDescData = file_vault_proto_rawDesc
)

func file_vault_proto_rawDescGZIP() []byte {
	file_vault_proto_rawDescOnce.Do(func() {
		file_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_proto_rawDescData)
	})
	return file_vault_proto_rawDescData
}

var file_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vault_proto_goTypes = []interface{}{
	(*Vault)(nil),         // 0: proto.Vault
	nil,                   // 1: proto.Vault.ContentsEntry
	(*ClientPayload)(nil), // 2: proto.ClientPayload
	(*ServerPayload)(nil), // 3: proto.ServerPayload
}
var file_vault_proto_depIdxs = []int32{
	1, // 0: proto.Vault.Contents:type_name -> proto.Vault.ContentsEntry
	2, // 1: proto.VaultServer.GetVault:input_type -> proto.ClientPayload
	2, // 2: proto.VaultServer.SetVault:input_type -> proto.ClientPayload
	3, // 3: proto.VaultServer.GetVault:output_type -> proto.ServerPayload
	3, // 4: proto.VaultServer.SetVault:output_type -> proto.ServerPayload
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vault_proto_init() }
func file_vault_proto_init() {
	if File_vault_proto != nil {
		return
	}
	file_srp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vault); i {
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
			RawDescriptor: file_vault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vault_proto_goTypes,
		DependencyIndexes: file_vault_proto_depIdxs,
		MessageInfos:      file_vault_proto_msgTypes,
	}.Build()
	File_vault_proto = out.File
	file_vault_proto_rawDesc = nil
	file_vault_proto_goTypes = nil
	file_vault_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VaultServerClient is the client API for VaultServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VaultServerClient interface {
	GetVault(ctx context.Context, opts ...grpc.CallOption) (VaultServer_GetVaultClient, error)
	SetVault(ctx context.Context, opts ...grpc.CallOption) (VaultServer_SetVaultClient, error)
}

type vaultServerClient struct {
	cc grpc.ClientConnInterface
}

func NewVaultServerClient(cc grpc.ClientConnInterface) VaultServerClient {
	return &vaultServerClient{cc}
}

func (c *vaultServerClient) GetVault(ctx context.Context, opts ...grpc.CallOption) (VaultServer_GetVaultClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VaultServer_serviceDesc.Streams[0], "/proto.VaultServer/GetVault", opts...)
	if err != nil {
		return nil, err
	}
	x := &vaultServerGetVaultClient{stream}
	return x, nil
}

type VaultServer_GetVaultClient interface {
	Send(*ClientPayload) error
	Recv() (*ServerPayload, error)
	grpc.ClientStream
}

type vaultServerGetVaultClient struct {
	grpc.ClientStream
}

func (x *vaultServerGetVaultClient) Send(m *ClientPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vaultServerGetVaultClient) Recv() (*ServerPayload, error) {
	m := new(ServerPayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vaultServerClient) SetVault(ctx context.Context, opts ...grpc.CallOption) (VaultServer_SetVaultClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VaultServer_serviceDesc.Streams[1], "/proto.VaultServer/SetVault", opts...)
	if err != nil {
		return nil, err
	}
	x := &vaultServerSetVaultClient{stream}
	return x, nil
}

type VaultServer_SetVaultClient interface {
	Send(*ClientPayload) error
	Recv() (*ServerPayload, error)
	grpc.ClientStream
}

type vaultServerSetVaultClient struct {
	grpc.ClientStream
}

func (x *vaultServerSetVaultClient) Send(m *ClientPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vaultServerSetVaultClient) Recv() (*ServerPayload, error) {
	m := new(ServerPayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VaultServerServer is the server API for VaultServer service.
type VaultServerServer interface {
	GetVault(VaultServer_GetVaultServer) error
	SetVault(VaultServer_SetVaultServer) error
}

// UnimplementedVaultServerServer can be embedded to have forward compatible implementations.
type UnimplementedVaultServerServer struct {
}

func (*UnimplementedVaultServerServer) GetVault(VaultServer_GetVaultServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVault not implemented")
}
func (*UnimplementedVaultServerServer) SetVault(VaultServer_SetVaultServer) error {
	return status.Errorf(codes.Unimplemented, "method SetVault not implemented")
}

func RegisterVaultServerServer(s *grpc.Server, srv VaultServerServer) {
	s.RegisterService(&_VaultServer_serviceDesc, srv)
}

func _VaultServer_GetVault_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VaultServerServer).GetVault(&vaultServerGetVaultServer{stream})
}

type VaultServer_GetVaultServer interface {
	Send(*ServerPayload) error
	Recv() (*ClientPayload, error)
	grpc.ServerStream
}

type vaultServerGetVaultServer struct {
	grpc.ServerStream
}

func (x *vaultServerGetVaultServer) Send(m *ServerPayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vaultServerGetVaultServer) Recv() (*ClientPayload, error) {
	m := new(ClientPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VaultServer_SetVault_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VaultServerServer).SetVault(&vaultServerSetVaultServer{stream})
}

type VaultServer_SetVaultServer interface {
	Send(*ServerPayload) error
	Recv() (*ClientPayload, error)
	grpc.ServerStream
}

type vaultServerSetVaultServer struct {
	grpc.ServerStream
}

func (x *vaultServerSetVaultServer) Send(m *ServerPayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vaultServerSetVaultServer) Recv() (*ClientPayload, error) {
	m := new(ClientPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VaultServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VaultServer",
	HandlerType: (*VaultServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetVault",
			Handler:       _VaultServer_GetVault_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SetVault",
			Handler:       _VaultServer_SetVault_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vault.proto",
}
