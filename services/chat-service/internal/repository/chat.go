package repository

import (
	"cloud.google.com/go/pubsub"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
)

type ChatRepository interface {
	CreateRoom(name string, members []string) (*pb.Room, error)
	JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error
	ListRoom(userID string) ([]*pb.Room, error)

	ListRoomMessages(roomID string, request *pb.ListRoomMessagesRequest) ([]*pb.Message, error)
	ListMembers(context context.Context, request *pb.ListMembersRequest) ([]*user.UserItem, error)

	SendMessage(msg *pb.Message, server pb.ChatService_SendMessageServer) error
}

type chatRepository struct {
	cache  cache.RoomCache
	pubsub *pubsub.Client
}

func NewChatRepository(cache cache.RoomCache, pubsub *pubsub.Client) ChatRepository {
	return &chatRepository{
		cache:  cache,
		pubsub: pubsub,
	}
}
