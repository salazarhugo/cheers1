// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: cheers/chat/v1/chat_service.proto

package party

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*Room, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (ChatService_JoinRoomClient, error)
	ListRoom(ctx context.Context, in *ListRoomRequest, opts ...grpc.CallOption) (ChatService_ListRoomClient, error)
	DeleteRoom(ctx context.Context, in *RoomId, opts ...grpc.CallOption) (*Empty, error)
	GetRoomId(ctx context.Context, in *GetRoomIdReq, opts ...grpc.CallOption) (*RoomId, error)
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
	LeaveRoom(ctx context.Context, in *RoomId, opts ...grpc.CallOption) (*Empty, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
	LikeMessage(ctx context.Context, in *LikeMessageReq, opts ...grpc.CallOption) (*Empty, error)
	UnlikeMessage(ctx context.Context, in *LikeMessageReq, opts ...grpc.CallOption) (*Empty, error)
	TypingChannel(ctx context.Context, opts ...grpc.CallOption) (ChatService_TypingChannelClient, error)
	TypingStart(ctx context.Context, in *TypingReq, opts ...grpc.CallOption) (*Empty, error)
	TypingEnd(ctx context.Context, in *TypingReq, opts ...grpc.CallOption) (*Empty, error)
	AddToken(ctx context.Context, in *AddTokenReq, opts ...grpc.CallOption) (*Empty, error)
	DeleteUser(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (ChatService_JoinRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/cheers.chat.v1.ChatService/JoinRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceJoinRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_JoinRoomClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceJoinRoomClient struct {
	grpc.ClientStream
}

func (x *chatServiceJoinRoomClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ListRoom(ctx context.Context, in *ListRoomRequest, opts ...grpc.CallOption) (ChatService_ListRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/cheers.chat.v1.ChatService/ListRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceListRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ListRoomClient interface {
	Recv() (*Room, error)
	grpc.ClientStream
}

type chatServiceListRoomClient struct {
	grpc.ClientStream
}

func (x *chatServiceListRoomClient) Recv() (*Room, error) {
	m := new(Room)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) DeleteRoom(ctx context.Context, in *RoomId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRoomId(ctx context.Context, in *GetRoomIdReq, opts ...grpc.CallOption) (*RoomId, error) {
	out := new(RoomId)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/GetRoomId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveRoom(ctx context.Context, in *RoomId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[2], "/cheers.chat.v1.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*SendMessageResponse, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) CloseAndRecv() (*SendMessageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) LikeMessage(ctx context.Context, in *LikeMessageReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/LikeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UnlikeMessage(ctx context.Context, in *LikeMessageReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/UnlikeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) TypingChannel(ctx context.Context, opts ...grpc.CallOption) (ChatService_TypingChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[3], "/cheers.chat.v1.ChatService/TypingChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceTypingChannelClient{stream}
	return x, nil
}

type ChatService_TypingChannelClient interface {
	Send(*TypingEvent) error
	Recv() (*TypingEvent, error)
	grpc.ClientStream
}

type chatServiceTypingChannelClient struct {
	grpc.ClientStream
}

func (x *chatServiceTypingChannelClient) Send(m *TypingEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceTypingChannelClient) Recv() (*TypingEvent, error) {
	m := new(TypingEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) TypingStart(ctx context.Context, in *TypingReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/TypingStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) TypingEnd(ctx context.Context, in *TypingReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/TypingEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddToken(ctx context.Context, in *AddTokenReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/AddToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteUser(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cheers.chat.v1.ChatService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	CreateChat(context.Context, *CreateChatReq) (*Room, error)
	JoinRoom(*JoinRoomRequest, ChatService_JoinRoomServer) error
	ListRoom(*ListRoomRequest, ChatService_ListRoomServer) error
	DeleteRoom(context.Context, *RoomId) (*Empty, error)
	GetRoomId(context.Context, *GetRoomIdReq) (*RoomId, error)
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	LeaveRoom(context.Context, *RoomId) (*Empty, error)
	SendMessage(ChatService_SendMessageServer) error
	LikeMessage(context.Context, *LikeMessageReq) (*Empty, error)
	UnlikeMessage(context.Context, *LikeMessageReq) (*Empty, error)
	TypingChannel(ChatService_TypingChannelServer) error
	TypingStart(context.Context, *TypingReq) (*Empty, error)
	TypingEnd(context.Context, *TypingReq) (*Empty, error)
	AddToken(context.Context, *AddTokenReq) (*Empty, error)
	DeleteUser(context.Context, *UserIdReq) (*Empty, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatReq) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) JoinRoom(*JoinRoomRequest, ChatService_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChatServiceServer) ListRoom(*ListRoomRequest, ChatService_ListRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRoom not implemented")
}
func (UnimplementedChatServiceServer) DeleteRoom(context.Context, *RoomId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedChatServiceServer) GetRoomId(context.Context, *GetRoomIdReq) (*RoomId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomId not implemented")
}
func (UnimplementedChatServiceServer) ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedChatServiceServer) LeaveRoom(context.Context, *RoomId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) LikeMessage(context.Context, *LikeMessageReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeMessage not implemented")
}
func (UnimplementedChatServiceServer) UnlikeMessage(context.Context, *LikeMessageReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikeMessage not implemented")
}
func (UnimplementedChatServiceServer) TypingChannel(ChatService_TypingChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method TypingChannel not implemented")
}
func (UnimplementedChatServiceServer) TypingStart(context.Context, *TypingReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TypingStart not implemented")
}
func (UnimplementedChatServiceServer) TypingEnd(context.Context, *TypingReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TypingEnd not implemented")
}
func (UnimplementedChatServiceServer) AddToken(context.Context, *AddTokenReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (UnimplementedChatServiceServer) DeleteUser(context.Context, *UserIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).JoinRoom(m, &chatServiceJoinRoomServer{stream})
}

type ChatService_JoinRoomServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceJoinRoomServer struct {
	grpc.ServerStream
}

func (x *chatServiceJoinRoomServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_ListRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ListRoom(m, &chatServiceListRoomServer{stream})
}

type ChatService_ListRoomServer interface {
	Send(*Room) error
	grpc.ServerStream
}

type chatServiceListRoomServer struct {
	grpc.ServerStream
}

func (x *chatServiceListRoomServer) Send(m *Room) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteRoom(ctx, req.(*RoomId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRoomId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetRoomId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/GetRoomId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetRoomId(ctx, req.(*GetRoomIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeaveRoom(ctx, req.(*RoomId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	SendAndClose(*SendMessageResponse) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) SendAndClose(m *SendMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_LikeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LikeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/LikeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LikeMessage(ctx, req.(*LikeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UnlikeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UnlikeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/UnlikeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UnlikeMessage(ctx, req.(*LikeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_TypingChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).TypingChannel(&chatServiceTypingChannelServer{stream})
}

type ChatService_TypingChannelServer interface {
	Send(*TypingEvent) error
	Recv() (*TypingEvent, error)
	grpc.ServerStream
}

type chatServiceTypingChannelServer struct {
	grpc.ServerStream
}

func (x *chatServiceTypingChannelServer) Send(m *TypingEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceTypingChannelServer) Recv() (*TypingEvent, error) {
	m := new(TypingEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_TypingStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).TypingStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/TypingStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).TypingStart(ctx, req.(*TypingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_TypingEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).TypingEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/TypingEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).TypingEnd(ctx, req.(*TypingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddToken(ctx, req.(*AddTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheers.chat.v1.ChatService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteUser(ctx, req.(*UserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheers.chat.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _ChatService_DeleteRoom_Handler,
		},
		{
			MethodName: "GetRoomId",
			Handler:    _ChatService_GetRoomId_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _ChatService_ListMembers_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _ChatService_LeaveRoom_Handler,
		},
		{
			MethodName: "LikeMessage",
			Handler:    _ChatService_LikeMessage_Handler,
		},
		{
			MethodName: "UnlikeMessage",
			Handler:    _ChatService_UnlikeMessage_Handler,
		},
		{
			MethodName: "TypingStart",
			Handler:    _ChatService_TypingStart_Handler,
		},
		{
			MethodName: "TypingEnd",
			Handler:    _ChatService_TypingEnd_Handler,
		},
		{
			MethodName: "AddToken",
			Handler:    _ChatService_AddToken_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ChatService_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinRoom",
			Handler:       _ChatService_JoinRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListRoom",
			Handler:       _ChatService_ListRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TypingChannel",
			Handler:       _ChatService_TypingChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cheers/chat/v1/chat_service.proto",
}
