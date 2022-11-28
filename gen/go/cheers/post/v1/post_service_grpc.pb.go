// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cheers/post/v1/post_service.proto

package post

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	// Create a new post
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List posts of a specific user
	ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostResponse, error)
	FeedPost(ctx context.Context, in *FeedPostRequest, opts ...grpc.CallOption) (*FeedPostResponse, error)
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error)
	UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*UnlikePostResponse, error)
	SavePost(ctx context.Context, in *SavePostRequest, opts ...grpc.CallOption) (*SavePostResponse, error)
	UnsavePost(ctx context.Context, in *UnsavePostRequest, opts ...grpc.CallOption) (*UnsavePostResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostResponse, error) {
	out := new(ListPostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/ListPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) FeedPost(ctx context.Context, in *FeedPostRequest, opts ...grpc.CallOption) (*FeedPostResponse, error) {
	out := new(FeedPostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/FeedPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error) {
	out := new(LikePostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*UnlikePostResponse, error) {
	out := new(UnlikePostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/UnlikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SavePost(ctx context.Context, in *SavePostRequest, opts ...grpc.CallOption) (*SavePostResponse, error) {
	out := new(SavePostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/SavePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UnsavePost(ctx context.Context, in *UnsavePostRequest, opts ...grpc.CallOption) (*UnsavePostResponse, error) {
	out := new(UnsavePostResponse)
	err := c.cc.Invoke(ctx, "/cheers.post.v1.PostService/UnsavePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	// Create a new post
	CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error)
	GetPost(context.Context, *GetPostRequest) (*PostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*PostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*emptypb.Empty, error)
	// List posts of a specific user
	ListPost(context.Context, *ListPostRequest) (*ListPostResponse, error)
	FeedPost(context.Context, *FeedPostRequest) (*FeedPostResponse, error)
	LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error)
	UnlikePost(context.Context, *UnlikePostRequest) (*UnlikePostResponse, error)
	SavePost(context.Context, *SavePostRequest) (*SavePostResponse, error)
	UnsavePost(context.Context, *UnsavePostRequest) (*UnsavePostResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *UpdatePostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) ListPost(context.Context, *ListPostRequest) (*ListPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedPostServiceServer) FeedPost(context.Context, *FeedPostRequest) (*FeedPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedPost not implemented")
}
func (UnimplementedPostServiceServer) LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedPostServiceServer) UnlikePost(context.Context, *UnlikePostRequest) (*UnlikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikePost not implemented")
}
func (UnimplementedPostServiceServer) SavePost(context.Context, *SavePostRequest) (*SavePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
}
func (UnimplementedPostServiceServer) UnsavePost(context.Context, *UnsavePostRequest) (*UnsavePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsavePost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/ListPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPost(ctx, req.(*ListPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_FeedPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).FeedPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/FeedPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).FeedPost(ctx, req.(*FeedPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UnlikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UnlikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/UnlikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UnlikePost(ctx, req.(*UnlikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/SavePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SavePost(ctx, req.(*SavePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UnsavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsavePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UnsavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.post.v1.PostService/UnsavePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UnsavePost(ctx, req.(*UnsavePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheers.post.v1.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "ListPost",
			Handler:    _PostService_ListPost_Handler,
		},
		{
			MethodName: "FeedPost",
			Handler:    _PostService_FeedPost_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _PostService_LikePost_Handler,
		},
		{
			MethodName: "UnlikePost",
			Handler:    _PostService_UnlikePost_Handler,
		},
		{
			MethodName: "SavePost",
			Handler:    _PostService_SavePost_Handler,
		},
		{
			MethodName: "UnsavePost",
			Handler:    _PostService_UnsavePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheers/post/v1/post_service.proto",
}
