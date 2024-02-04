// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cheers/friendship/v1/friendship_service.proto

package friendship

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FriendshipServiceClient is the client API for FriendshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendshipServiceClient interface {
	//
	//Send a friend request
	CreateFriendRequest(ctx context.Context, in *CreateFriendRequestRequest, opts ...grpc.CallOption) (*CreateFriendRequestResponse, error)
	//
	//Accept a friend request
	AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error)
	//
	//Get friend list of a specific user
	ListFriend(ctx context.Context, in *ListFriendRequest, opts ...grpc.CallOption) (*ListFriendResponse, error)
	//
	//Get friend list of a specific user
	ListFriendIds(ctx context.Context, in *ListFriendIdsRequest, opts ...grpc.CallOption) (*ListFriendIdsResponse, error)
	//
	//Get friend requests list of a specific user
	ListFriendRequests(ctx context.Context, in *ListFriendRequestsRequest, opts ...grpc.CallOption) (*ListFriendRequestsResponse, error)
	//
	//Refuse a friend request
	DeleteFriendRequest(ctx context.Context, in *DeleteFriendRequestRequest, opts ...grpc.CallOption) (*DeleteFriendRequestResponse, error)
	//
	//Delete a friend
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest2, opts ...grpc.CallOption) (*DeleteFriendResponse, error)
}

type friendshipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendshipServiceClient(cc grpc.ClientConnInterface) FriendshipServiceClient {
	return &friendshipServiceClient{cc}
}

func (c *friendshipServiceClient) CreateFriendRequest(ctx context.Context, in *CreateFriendRequestRequest, opts ...grpc.CallOption) (*CreateFriendRequestResponse, error) {
	out := new(CreateFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/CreateFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error) {
	out := new(AcceptFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/AcceptFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) ListFriend(ctx context.Context, in *ListFriendRequest, opts ...grpc.CallOption) (*ListFriendResponse, error) {
	out := new(ListFriendResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/ListFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) ListFriendIds(ctx context.Context, in *ListFriendIdsRequest, opts ...grpc.CallOption) (*ListFriendIdsResponse, error) {
	out := new(ListFriendIdsResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/ListFriendIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) ListFriendRequests(ctx context.Context, in *ListFriendRequestsRequest, opts ...grpc.CallOption) (*ListFriendRequestsResponse, error) {
	out := new(ListFriendRequestsResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/ListFriendRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) DeleteFriendRequest(ctx context.Context, in *DeleteFriendRequestRequest, opts ...grpc.CallOption) (*DeleteFriendRequestResponse, error) {
	out := new(DeleteFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/DeleteFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest2, opts ...grpc.CallOption) (*DeleteFriendResponse, error) {
	out := new(DeleteFriendResponse)
	err := c.cc.Invoke(ctx, "/cheers.friendship.v1.FriendshipService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendshipServiceServer is the server API for FriendshipService service.
// All implementations must embed UnimplementedFriendshipServiceServer
// for forward compatibility
type FriendshipServiceServer interface {
	//
	//Send a friend request
	CreateFriendRequest(context.Context, *CreateFriendRequestRequest) (*CreateFriendRequestResponse, error)
	//
	//Accept a friend request
	AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error)
	//
	//Get friend list of a specific user
	ListFriend(context.Context, *ListFriendRequest) (*ListFriendResponse, error)
	//
	//Get friend list of a specific user
	ListFriendIds(context.Context, *ListFriendIdsRequest) (*ListFriendIdsResponse, error)
	//
	//Get friend requests list of a specific user
	ListFriendRequests(context.Context, *ListFriendRequestsRequest) (*ListFriendRequestsResponse, error)
	//
	//Refuse a friend request
	DeleteFriendRequest(context.Context, *DeleteFriendRequestRequest) (*DeleteFriendRequestResponse, error)
	//
	//Delete a friend
	DeleteFriend(context.Context, *DeleteFriendRequest2) (*DeleteFriendResponse, error)
	mustEmbedUnimplementedFriendshipServiceServer()
}

// UnimplementedFriendshipServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendshipServiceServer struct {
}

func (UnimplementedFriendshipServiceServer) CreateFriendRequest(context.Context, *CreateFriendRequestRequest) (*CreateFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFriendRequest not implemented")
}
func (UnimplementedFriendshipServiceServer) AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriendRequest not implemented")
}
func (UnimplementedFriendshipServiceServer) ListFriend(context.Context, *ListFriendRequest) (*ListFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriend not implemented")
}
func (UnimplementedFriendshipServiceServer) ListFriendIds(context.Context, *ListFriendIdsRequest) (*ListFriendIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriendIds not implemented")
}
func (UnimplementedFriendshipServiceServer) ListFriendRequests(context.Context, *ListFriendRequestsRequest) (*ListFriendRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriendRequests not implemented")
}
func (UnimplementedFriendshipServiceServer) DeleteFriendRequest(context.Context, *DeleteFriendRequestRequest) (*DeleteFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriendRequest not implemented")
}
func (UnimplementedFriendshipServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest2) (*DeleteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendshipServiceServer) mustEmbedUnimplementedFriendshipServiceServer() {}

// UnsafeFriendshipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendshipServiceServer will
// result in compilation errors.
type UnsafeFriendshipServiceServer interface {
	mustEmbedUnimplementedFriendshipServiceServer()
}

func RegisterFriendshipServiceServer(s grpc.ServiceRegistrar, srv FriendshipServiceServer) {
	s.RegisterService(&FriendshipService_ServiceDesc, srv)
}

func _FriendshipService_CreateFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).CreateFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/CreateFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).CreateFriendRequest(ctx, req.(*CreateFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_AcceptFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).AcceptFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/AcceptFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).AcceptFriendRequest(ctx, req.(*AcceptFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_ListFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).ListFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/ListFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).ListFriend(ctx, req.(*ListFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_ListFriendIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).ListFriendIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/ListFriendIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).ListFriendIds(ctx, req.(*ListFriendIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_ListFriendRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).ListFriendRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/ListFriendRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).ListFriendRequests(ctx, req.(*ListFriendRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_DeleteFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).DeleteFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/DeleteFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).DeleteFriendRequest(ctx, req.(*DeleteFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.friendship.v1.FriendshipService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendshipService_ServiceDesc is the grpc.ServiceDesc for FriendshipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendshipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheers.friendship.v1.FriendshipService",
	HandlerType: (*FriendshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFriendRequest",
			Handler:    _FriendshipService_CreateFriendRequest_Handler,
		},
		{
			MethodName: "AcceptFriendRequest",
			Handler:    _FriendshipService_AcceptFriendRequest_Handler,
		},
		{
			MethodName: "ListFriend",
			Handler:    _FriendshipService_ListFriend_Handler,
		},
		{
			MethodName: "ListFriendIds",
			Handler:    _FriendshipService_ListFriendIds_Handler,
		},
		{
			MethodName: "ListFriendRequests",
			Handler:    _FriendshipService_ListFriendRequests_Handler,
		},
		{
			MethodName: "DeleteFriendRequest",
			Handler:    _FriendshipService_DeleteFriendRequest_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _FriendshipService_DeleteFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheers/friendship/v1/friendship_service.proto",
}
