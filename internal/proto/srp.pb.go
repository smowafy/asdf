// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: srp.proto

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

type ClientPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *ClientPayload) Reset() {
	*x = ClientPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPayload) ProtoMessage() {}

func (x *ClientPayload) ProtoReflect() protoreflect.Message {
	mi := &file_srp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPayload.ProtoReflect.Descriptor instead.
func (*ClientPayload) Descriptor() ([]byte, []int) {
	return file_srp_proto_rawDescGZIP(), []int{0}
}

func (x *ClientPayload) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ServerPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *ServerPayload) Reset() {
	*x = ServerPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPayload) ProtoMessage() {}

func (x *ServerPayload) ProtoReflect() protoreflect.Message {
	mi := &file_srp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPayload.ProtoReflect.Descriptor instead.
func (*ServerPayload) Descriptor() ([]byte, []int) {
	return file_srp_proto_rawDescGZIP(), []int{1}
}

func (x *ServerPayload) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ClientVerifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Verif     string `protobuf:"bytes,2,opt,name=Verif,proto3" json:"Verif,omitempty"`
	AccountId string `protobuf:"bytes,3,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
}

func (x *ClientVerifier) Reset() {
	*x = ClientVerifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientVerifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientVerifier) ProtoMessage() {}

func (x *ClientVerifier) ProtoReflect() protoreflect.Message {
	mi := &file_srp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientVerifier.ProtoReflect.Descriptor instead.
func (*ClientVerifier) Descriptor() ([]byte, []int) {
	return file_srp_proto_rawDescGZIP(), []int{2}
}

func (x *ClientVerifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientVerifier) GetVerif() string {
	if x != nil {
		return x.Verif
	}
	return ""
}

func (x *ClientVerifier) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type VerifierStored struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *VerifierStored) Reset() {
	*x = VerifierStored{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifierStored) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifierStored) ProtoMessage() {}

func (x *VerifierStored) ProtoReflect() protoreflect.Message {
	mi := &file_srp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifierStored.ProtoReflect.Descriptor instead.
func (*VerifierStored) Descriptor() ([]byte, []int) {
	return file_srp_proto_rawDescGZIP(), []int{3}
}

func (x *VerifierStored) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_srp_proto protoreflect.FileDescriptor

var file_srp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x54, 0x0a, 0x0e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x65, 0x72, 0x69, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x45,
	0x0a, 0x09, 0x53, 0x72, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_srp_proto_rawDescOnce sync.Once
	file_srp_proto_rawDescData = file_srp_proto_rawDesc
)

func file_srp_proto_rawDescGZIP() []byte {
	file_srp_proto_rawDescOnce.Do(func() {
		file_srp_proto_rawDescData = protoimpl.X.CompressGZIP(file_srp_proto_rawDescData)
	})
	return file_srp_proto_rawDescData
}

var file_srp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_srp_proto_goTypes = []interface{}{
	(*ClientPayload)(nil),  // 0: proto.ClientPayload
	(*ServerPayload)(nil),  // 1: proto.ServerPayload
	(*ClientVerifier)(nil), // 2: proto.ClientVerifier
	(*VerifierStored)(nil), // 3: proto.VerifierStored
}
var file_srp_proto_depIdxs = []int32{
	2, // 0: proto.SrpServer.SignUp:input_type -> proto.ClientVerifier
	3, // 1: proto.SrpServer.SignUp:output_type -> proto.VerifierStored
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_srp_proto_init() }
func file_srp_proto_init() {
	if File_srp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_srp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPayload); i {
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
		file_srp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPayload); i {
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
		file_srp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientVerifier); i {
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
		file_srp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifierStored); i {
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
			RawDescriptor: file_srp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_srp_proto_goTypes,
		DependencyIndexes: file_srp_proto_depIdxs,
		MessageInfos:      file_srp_proto_msgTypes,
	}.Build()
	File_srp_proto = out.File
	file_srp_proto_rawDesc = nil
	file_srp_proto_goTypes = nil
	file_srp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SrpServerClient is the client API for SrpServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SrpServerClient interface {
	SignUp(ctx context.Context, in *ClientVerifier, opts ...grpc.CallOption) (*VerifierStored, error)
}

type srpServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSrpServerClient(cc grpc.ClientConnInterface) SrpServerClient {
	return &srpServerClient{cc}
}

func (c *srpServerClient) SignUp(ctx context.Context, in *ClientVerifier, opts ...grpc.CallOption) (*VerifierStored, error) {
	out := new(VerifierStored)
	err := c.cc.Invoke(ctx, "/proto.SrpServer/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrpServerServer is the server API for SrpServer service.
type SrpServerServer interface {
	SignUp(context.Context, *ClientVerifier) (*VerifierStored, error)
}

// UnimplementedSrpServerServer can be embedded to have forward compatible implementations.
type UnimplementedSrpServerServer struct {
}

func (*UnimplementedSrpServerServer) SignUp(context.Context, *ClientVerifier) (*VerifierStored, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func RegisterSrpServerServer(s *grpc.Server, srv SrpServerServer) {
	s.RegisterService(&_SrpServer_serviceDesc, srv)
}

func _SrpServer_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientVerifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrpServerServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SrpServer/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrpServerServer).SignUp(ctx, req.(*ClientVerifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _SrpServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SrpServer",
	HandlerType: (*SrpServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _SrpServer_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srp.proto",
}