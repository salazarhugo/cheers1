// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: cheers/api/v1/main.proto

package proto

import (
	context "context"
	proto "github.com/salazarhugo/cheers1/genproto/cheers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MainClient is the client API for Main service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*proto.Party, error)
	DeleteParty(ctx context.Context, in *DeletePartyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mainClient struct {
	cc grpc.ClientConnInterface
}

func NewMainClient(cc grpc.ClientConnInterface) MainClient {
	return &mainClient{cc}
}

func (c *mainClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/cheers.api.v1.Main/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*proto.Party, error) {
	out := new(proto.Party)
	err := c.cc.Invoke(ctx, "/cheers.api.v1.Main/CreateParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeleteParty(ctx context.Context, in *DeletePartyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cheers.api.v1.Main/DeleteParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServer is the server API for Main service.
// All implementations must embed UnimplementedMainServer
// for forward compatibility
type MainServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateParty(context.Context, *CreatePartyRequest) (*proto.Party, error)
	DeleteParty(context.Context, *DeletePartyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMainServer()
}

// UnimplementedMainServer must be embedded to have forward compatible implementations.
type UnimplementedMainServer struct {
}

func (UnimplementedMainServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedMainServer) CreateParty(context.Context, *CreatePartyRequest) (*proto.Party, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParty not implemented")
}
func (UnimplementedMainServer) DeleteParty(context.Context, *DeletePartyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParty not implemented")
}
func (UnimplementedMainServer) mustEmbedUnimplementedMainServer() {}

// UnsafeMainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServer will
// result in compilation errors.
type UnsafeMainServer interface {
	mustEmbedUnimplementedMainServer()
}

func RegisterMainServer(s grpc.ServiceRegistrar, srv MainServer) {
	s.RegisterService(&Main_ServiceDesc, srv)
}

func _Main_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.api.v1.Main/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CreateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CreateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.api.v1.Main/CreateParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CreateParty(ctx, req.(*CreatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeleteParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.api.v1.Main/DeleteParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteParty(ctx, req.(*DeletePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Main_ServiceDesc is the grpc.ServiceDesc for Main service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Main_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheers.api.v1.Main",
	HandlerType: (*MainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Main_GetUser_Handler,
		},
		{
			MethodName: "CreateParty",
			Handler:    _Main_CreateParty_Handler,
		},
		{
			MethodName: "DeleteParty",
			Handler:    _Main_DeleteParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheers/api/v1/main.proto",
}
