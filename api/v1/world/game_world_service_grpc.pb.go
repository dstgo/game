// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/v1/world/game_world_service.proto

package world

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	World_Hello_FullMethodName = "/api.v1.world.World/Hello"
)

// WorldClient is the client API for World service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorldClient interface {
	Hello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type worldClient struct {
	cc grpc.ClientConnInterface
}

func NewWorldClient(cc grpc.ClientConnInterface) WorldClient {
	return &worldClient{cc}
}

func (c *worldClient) Hello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, World_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorldServer is the server API for World service.
// All implementations must embed UnimplementedWorldServer
// for forward compatibility.
type WorldServer interface {
	Hello(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedWorldServer()
}

// UnimplementedWorldServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorldServer struct{}

func (UnimplementedWorldServer) Hello(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedWorldServer) mustEmbedUnimplementedWorldServer() {}
func (UnimplementedWorldServer) testEmbeddedByValue()               {}

// UnsafeWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorldServer will
// result in compilation errors.
type UnsafeWorldServer interface {
	mustEmbedUnimplementedWorldServer()
}

func RegisterWorldServer(s grpc.ServiceRegistrar, srv WorldServer) {
	// If the following call pancis, it indicates UnimplementedWorldServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&World_ServiceDesc, srv)
}

func _World_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: World_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldServer).Hello(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// World_ServiceDesc is the grpc.ServiceDesc for World service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var World_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.world.World",
	HandlerType: (*WorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _World_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/world/game_world_service.proto",
}
