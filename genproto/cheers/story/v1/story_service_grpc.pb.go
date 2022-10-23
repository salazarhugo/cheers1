// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: cheers/story/v1/story_service.proto

package story

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

// StoryServiceClient is the client API for StoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoryServiceClient interface {
	CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error)
	GetStory(ctx context.Context, in *GetStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error)
	UpdateStory(ctx context.Context, in *UpdateStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error)
	DeleteStory(ctx context.Context, in *DeleteStoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FeedStory(ctx context.Context, in *FeedStoryRequest, opts ...grpc.CallOption) (*FeedStoryResponse, error)
	//
	//Return stories of a specific user
	ListUserStory(ctx context.Context, in *ListUserStoryRequest, opts ...grpc.CallOption) (*ListUserStoryResponse, error)
	ViewStory(ctx context.Context, in *ViewStoryRequest, opts ...grpc.CallOption) (*ViewStoryResponse, error)
	LikeStory(ctx context.Context, in *LikeStoryRequest, opts ...grpc.CallOption) (*LikeStoryResponse, error)
	UnlikeStory(ctx context.Context, in *UnlikeStoryRequest, opts ...grpc.CallOption) (*UnlikeStoryResponse, error)
	SaveStory(ctx context.Context, in *SaveStoryRequest, opts ...grpc.CallOption) (*SaveStoryResponse, error)
	UnsaveStory(ctx context.Context, in *UnsaveStoryRequest, opts ...grpc.CallOption) (*UnsaveStoryResponse, error)
}

type storyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoryServiceClient(cc grpc.ClientConnInterface) StoryServiceClient {
	return &storyServiceClient{cc}
}

func (c *storyServiceClient) CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error) {
	out := new(StoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/CreateStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) GetStory(ctx context.Context, in *GetStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error) {
	out := new(StoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/GetStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) UpdateStory(ctx context.Context, in *UpdateStoryRequest, opts ...grpc.CallOption) (*StoryResponse, error) {
	out := new(StoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/UpdateStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) DeleteStory(ctx context.Context, in *DeleteStoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/DeleteStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) FeedStory(ctx context.Context, in *FeedStoryRequest, opts ...grpc.CallOption) (*FeedStoryResponse, error) {
	out := new(FeedStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/FeedStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) ListUserStory(ctx context.Context, in *ListUserStoryRequest, opts ...grpc.CallOption) (*ListUserStoryResponse, error) {
	out := new(ListUserStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/ListUserStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) ViewStory(ctx context.Context, in *ViewStoryRequest, opts ...grpc.CallOption) (*ViewStoryResponse, error) {
	out := new(ViewStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/ViewStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) LikeStory(ctx context.Context, in *LikeStoryRequest, opts ...grpc.CallOption) (*LikeStoryResponse, error) {
	out := new(LikeStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/LikeStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) UnlikeStory(ctx context.Context, in *UnlikeStoryRequest, opts ...grpc.CallOption) (*UnlikeStoryResponse, error) {
	out := new(UnlikeStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/UnlikeStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) SaveStory(ctx context.Context, in *SaveStoryRequest, opts ...grpc.CallOption) (*SaveStoryResponse, error) {
	out := new(SaveStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/SaveStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) UnsaveStory(ctx context.Context, in *UnsaveStoryRequest, opts ...grpc.CallOption) (*UnsaveStoryResponse, error) {
	out := new(UnsaveStoryResponse)
	err := c.cc.Invoke(ctx, "/cheers.story.v1.StoryService/UnsaveStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoryServiceServer is the server API for StoryService service.
// All implementations must embed UnimplementedStoryServiceServer
// for forward compatibility
type StoryServiceServer interface {
	CreateStory(context.Context, *CreateStoryRequest) (*StoryResponse, error)
	GetStory(context.Context, *GetStoryRequest) (*StoryResponse, error)
	UpdateStory(context.Context, *UpdateStoryRequest) (*StoryResponse, error)
	DeleteStory(context.Context, *DeleteStoryRequest) (*emptypb.Empty, error)
	FeedStory(context.Context, *FeedStoryRequest) (*FeedStoryResponse, error)
	//
	//Return stories of a specific user
	ListUserStory(context.Context, *ListUserStoryRequest) (*ListUserStoryResponse, error)
	ViewStory(context.Context, *ViewStoryRequest) (*ViewStoryResponse, error)
	LikeStory(context.Context, *LikeStoryRequest) (*LikeStoryResponse, error)
	UnlikeStory(context.Context, *UnlikeStoryRequest) (*UnlikeStoryResponse, error)
	SaveStory(context.Context, *SaveStoryRequest) (*SaveStoryResponse, error)
	UnsaveStory(context.Context, *UnsaveStoryRequest) (*UnsaveStoryResponse, error)
	mustEmbedUnimplementedStoryServiceServer()
}

// UnimplementedStoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoryServiceServer struct {
}

func (UnimplementedStoryServiceServer) CreateStory(context.Context, *CreateStoryRequest) (*StoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStory not implemented")
}
func (UnimplementedStoryServiceServer) GetStory(context.Context, *GetStoryRequest) (*StoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStory not implemented")
}
func (UnimplementedStoryServiceServer) UpdateStory(context.Context, *UpdateStoryRequest) (*StoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStory not implemented")
}
func (UnimplementedStoryServiceServer) DeleteStory(context.Context, *DeleteStoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStory not implemented")
}
func (UnimplementedStoryServiceServer) FeedStory(context.Context, *FeedStoryRequest) (*FeedStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedStory not implemented")
}
func (UnimplementedStoryServiceServer) ListUserStory(context.Context, *ListUserStoryRequest) (*ListUserStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserStory not implemented")
}
func (UnimplementedStoryServiceServer) ViewStory(context.Context, *ViewStoryRequest) (*ViewStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewStory not implemented")
}
func (UnimplementedStoryServiceServer) LikeStory(context.Context, *LikeStoryRequest) (*LikeStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeStory not implemented")
}
func (UnimplementedStoryServiceServer) UnlikeStory(context.Context, *UnlikeStoryRequest) (*UnlikeStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikeStory not implemented")
}
func (UnimplementedStoryServiceServer) SaveStory(context.Context, *SaveStoryRequest) (*SaveStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveStory not implemented")
}
func (UnimplementedStoryServiceServer) UnsaveStory(context.Context, *UnsaveStoryRequest) (*UnsaveStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsaveStory not implemented")
}
func (UnimplementedStoryServiceServer) mustEmbedUnimplementedStoryServiceServer() {}

// UnsafeStoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoryServiceServer will
// result in compilation errors.
type UnsafeStoryServiceServer interface {
	mustEmbedUnimplementedStoryServiceServer()
}

func RegisterStoryServiceServer(s grpc.ServiceRegistrar, srv StoryServiceServer) {
	s.RegisterService(&StoryService_ServiceDesc, srv)
}

func _StoryService_CreateStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).CreateStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/CreateStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).CreateStory(ctx, req.(*CreateStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_GetStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).GetStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/GetStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).GetStory(ctx, req.(*GetStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_UpdateStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).UpdateStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/UpdateStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).UpdateStory(ctx, req.(*UpdateStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_DeleteStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).DeleteStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/DeleteStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).DeleteStory(ctx, req.(*DeleteStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_FeedStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).FeedStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/FeedStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).FeedStory(ctx, req.(*FeedStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_ListUserStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).ListUserStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/ListUserStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).ListUserStory(ctx, req.(*ListUserStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_ViewStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).ViewStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/ViewStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).ViewStory(ctx, req.(*ViewStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_LikeStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).LikeStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/LikeStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).LikeStory(ctx, req.(*LikeStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_UnlikeStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlikeStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).UnlikeStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/UnlikeStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).UnlikeStory(ctx, req.(*UnlikeStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_SaveStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).SaveStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/SaveStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).SaveStory(ctx, req.(*SaveStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_UnsaveStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsaveStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).UnsaveStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.story.v1.StoryService/UnsaveStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).UnsaveStory(ctx, req.(*UnsaveStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoryService_ServiceDesc is the grpc.ServiceDesc for StoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheers.story.v1.StoryService",
	HandlerType: (*StoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStory",
			Handler:    _StoryService_CreateStory_Handler,
		},
		{
			MethodName: "GetStory",
			Handler:    _StoryService_GetStory_Handler,
		},
		{
			MethodName: "UpdateStory",
			Handler:    _StoryService_UpdateStory_Handler,
		},
		{
			MethodName: "DeleteStory",
			Handler:    _StoryService_DeleteStory_Handler,
		},
		{
			MethodName: "FeedStory",
			Handler:    _StoryService_FeedStory_Handler,
		},
		{
			MethodName: "ListUserStory",
			Handler:    _StoryService_ListUserStory_Handler,
		},
		{
			MethodName: "ViewStory",
			Handler:    _StoryService_ViewStory_Handler,
		},
		{
			MethodName: "LikeStory",
			Handler:    _StoryService_LikeStory_Handler,
		},
		{
			MethodName: "UnlikeStory",
			Handler:    _StoryService_UnlikeStory_Handler,
		},
		{
			MethodName: "SaveStory",
			Handler:    _StoryService_SaveStory_Handler,
		},
		{
			MethodName: "UnsaveStory",
			Handler:    _StoryService_UnsaveStory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheers/story/v1/story_service.proto",
}
