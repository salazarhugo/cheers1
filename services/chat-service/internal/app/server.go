package app

import (
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"google.golang.org/api/chat/v1"
	"sync"
	"time"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	mu             sync.Mutex
	chatRepository repository.ChatRepository
	channel        map[string][]chan *chat.Message
}

func NewServer() *Server {
	chatRepository := repository.NewChatRepository(
		time.Duration(time.Duration.Hours(1)),
		redis.NewClient(&redis.Options{
			Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
			Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
			DB:       0,
		}),
	)
	return &Server{
		UnimplementedChatServiceServer: pb.UnimplementedChatServiceServer{},
		mu:                             sync.Mutex{},
		chatRepository:                 chatRepository,
		channel:                        make(map[string][]chan *chat.Message),
	}
}

func (s *Server) GetRoomId(ctx context.Context, req *pb.GetRoomIdReq) (*pb.RoomId, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetRoomMembers(ctx context.Context, id *pb.RoomId) (*pb.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) AddToken(ctx context.Context, req *pb.AddTokenReq) (*pb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.UserIdReq) (*pb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	panic("implement me")
}
