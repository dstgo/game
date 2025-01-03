// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/v1/mod/game_mod_service.proto

package mod

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
	Mod_Hello_FullMethodName = "/api.v1.mod.Mod/Hello"
)

// ModClient is the client API for Mod service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModClient interface {
	Hello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type modClient struct {
	cc grpc.ClientConnInterface
}

func NewModClient(cc grpc.ClientConnInterface) ModClient {
	return &modClient{cc}
}

func (c *modClient) Hello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Mod_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModServer is the server API for Mod service.
// All implementations must embed UnimplementedModServer
// for forward compatibility.
type ModServer interface {
	Hello(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedModServer()
}

// UnimplementedModServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModServer struct{}

func (UnimplementedModServer) Hello(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedModServer) mustEmbedUnimplementedModServer() {}
func (UnimplementedModServer) testEmbeddedByValue()             {}

// UnsafeModServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModServer will
// result in compilation errors.
type UnsafeModServer interface {
	mustEmbedUnimplementedModServer()
}

func RegisterModServer(s grpc.ServiceRegistrar, srv ModServer) {
	// If the following call pancis, it indicates UnimplementedModServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mod_ServiceDesc, srv)
}

func _Mod_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mod_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModServer).Hello(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Mod_ServiceDesc is the grpc.ServiceDesc for Mod service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mod_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.mod.Mod",
	HandlerType: (*ModServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Mod_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/mod/game_mod_service.proto",
}
