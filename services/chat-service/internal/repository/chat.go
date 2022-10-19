package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
)

type ChatRepository interface {
	SendMessage(msg *pb.Message, server pb.ChatService_SendMessageServer) error
	JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error
	CreateRoom(name string, members []string) (*pb.Room, error)
}

type chatRepository struct {
	cache cache.RoomCache
}

func NewChatRepository(cache cache.RoomCache) ChatRepository {
	return &chatRepository{
		cache: cache,
	}
}
