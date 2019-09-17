// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v3alpha/sds.proto

package envoy_service_discovery_v3alpha

import (
	context "context"
	fmt "fmt"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1969cf11de54a1fa, []int{0}
}

func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (m *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(m, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.discovery.v3alpha.SdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v3alpha/sds.proto", fileDescriptor_1969cf11de54a1fa)
}

var fileDescriptor_1969cf11de54a1fa = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0x87, 0x4d, 0x17, 0x22, 0x43, 0xdd, 0x64, 0xe1, 0x22, 0x08, 0xc5, 0xa8, 0xd0, 0x0a, 0x4e,
	0xc4, 0xee, 0xea, 0xae, 0x04, 0xd7, 0xc5, 0x5c, 0xc0, 0x31, 0x79, 0xb4, 0x03, 0xc9, 0xbc, 0x71,
	0xde, 0x34, 0x98, 0x9d, 0x78, 0x05, 0x8f, 0xe6, 0x15, 0x5c, 0x78, 0x0c, 0xe9, 0x4c, 0x12, 0xf1,
	0x0f, 0xea, 0xa6, 0xeb, 0xf9, 0xe6, 0xf7, 0x3d, 0xf8, 0xd8, 0x04, 0x54, 0x8d, 0x4d, 0x42, 0x60,
	0x6a, 0x99, 0x43, 0x52, 0x48, 0xca, 0xb1, 0x06, 0xd3, 0x24, 0xf5, 0x54, 0x94, 0x7a, 0x25, 0x12,
	0x2a, 0x88, 0x6b, 0x83, 0x16, 0xc3, 0x91, 0x43, 0x79, 0x8b, 0xf2, 0x1e, 0xe5, 0x2d, 0x1a, 0x1d,
	0xf9, 0x2d, 0xa1, 0x65, 0xff, 0xfb, 0x03, 0x72, 0x1b, 0xd1, 0xe1, 0x12, 0x71, 0x59, 0x82, 0x63,
	0x84, 0x52, 0x68, 0x85, 0x95, 0xa8, 0x5a, 0x43, 0xcc, 0xd8, 0x5e, 0x56, 0x50, 0xba, 0xae, 0xaa,
	0xe6, 0xf2, 0x6d, 0xc0, 0x0e, 0x32, 0xc8, 0x0d, 0xd8, 0xb4, 0xdb, 0xc8, 0xbc, 0x39, 0x94, 0x6c,
	0x98, 0x42, 0x69, 0x85, 0x7f, 0xa6, 0x70, 0xcc, 0xfd, 0x65, 0x42, 0xcb, 0xee, 0x16, 0xee, 0x80,
	0xfe, 0xe7, 0x0d, 0xdc, 0xaf, 0x81, 0x6c, 0x34, 0xf9, 0x07, 0x49, 0x1a, 0x15, 0x41, 0xbc, 0x33,
	0x0e, 0x2e, 0x82, 0xf0, 0x96, 0xed, 0x67, 0xd6, 0x80, 0xa8, 0x3a, 0xd7, 0xf1, 0x4f, 0x0b, 0x5f,
	0x35, 0x27, 0xbf, 0x43, 0x9f, 0x0c, 0x8f, 0x01, 0x1b, 0x5e, 0x83, 0xcd, 0x57, 0x5b, 0x30, 0x9c,
	0x3e, 0xbd, 0xbc, 0x3e, 0x0f, 0x46, 0x71, 0xf4, 0xbd, 0xc7, 0x8c, 0xbc, 0x6e, 0x16, 0x9c, 0xcd,
	0xaf, 0xd8, 0xb9, 0x44, 0x3f, 0xa8, 0x0d, 0x3e, 0x34, 0xfc, 0x8f, 0xd0, 0xf3, 0x4d, 0xa5, 0xc5,
	0xa6, 0xd8, 0x22, 0xb8, 0xdb, 0x75, 0xe9, 0xa6, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x0f,
	0x12, 0x07, 0x49, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error)
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v3alpha.SecretDiscoveryService/DeltaSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceDeltaSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_DeltaSecretsClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceDeltaSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v3alpha.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v3alpha.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	DeltaSecrets(SecretDiscoveryService_DeltaSecretsServer) error
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_DeltaSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).DeltaSecrets(&secretDiscoveryServiceDeltaSecretsServer{stream})
}

type SecretDiscoveryService_DeltaSecretsServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceDeltaSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v3alpha.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3alpha.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaSecrets",
			Handler:       _SecretDiscoveryService_DeltaSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v3alpha/sds.proto",
}