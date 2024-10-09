// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: bidi.proto

package bidipb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BidiService_BidiChat_FullMethodName = "/bidiservice.BidiService/BidiChat"
)

// BidiServiceClient is the client API for BidiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition with a bidirectional streaming RPC
type BidiServiceClient interface {
	BidiChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BidiRequest, BidiResponse], error)
}

type bidiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBidiServiceClient(cc grpc.ClientConnInterface) BidiServiceClient {
	return &bidiServiceClient{cc}
}

func (c *bidiServiceClient) BidiChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BidiRequest, BidiResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BidiService_ServiceDesc.Streams[0], BidiService_BidiChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BidiRequest, BidiResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BidiService_BidiChatClient = grpc.BidiStreamingClient[BidiRequest, BidiResponse]

// BidiServiceServer is the server API for BidiService service.
// All implementations must embed UnimplementedBidiServiceServer
// for forward compatibility.
//
// Service definition with a bidirectional streaming RPC
type BidiServiceServer interface {
	BidiChat(grpc.BidiStreamingServer[BidiRequest, BidiResponse]) error
	mustEmbedUnimplementedBidiServiceServer()
}

// UnimplementedBidiServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBidiServiceServer struct{}

func (UnimplementedBidiServiceServer) BidiChat(grpc.BidiStreamingServer[BidiRequest, BidiResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BidiChat not implemented")
}
func (UnimplementedBidiServiceServer) mustEmbedUnimplementedBidiServiceServer() {}
func (UnimplementedBidiServiceServer) testEmbeddedByValue()                     {}

// UnsafeBidiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidiServiceServer will
// result in compilation errors.
type UnsafeBidiServiceServer interface {
	mustEmbedUnimplementedBidiServiceServer()
}

func RegisterBidiServiceServer(s grpc.ServiceRegistrar, srv BidiServiceServer) {
	// If the following call pancis, it indicates UnimplementedBidiServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BidiService_ServiceDesc, srv)
}

func _BidiService_BidiChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidiServiceServer).BidiChat(&grpc.GenericServerStream[BidiRequest, BidiResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BidiService_BidiChatServer = grpc.BidiStreamingServer[BidiRequest, BidiResponse]

// BidiService_ServiceDesc is the grpc.ServiceDesc for BidiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bidiservice.BidiService",
	HandlerType: (*BidiServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidiChat",
			Handler:       _BidiService_BidiChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bidi.proto",
}
