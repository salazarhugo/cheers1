package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"google.golang.org/api/chat/v1"
	"sync"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	mu             sync.Mutex
	chatRepository repository.ChatRepository
	channel        map[string][]chan *chat.Message
}

func NewServer() *Server {
	chatRepository := repository.NewChatRepository()

	return &Server{
		UnimplementedChatServiceServer: pb.UnimplementedChatServiceServer{},
		mu:                             sync.Mutex{},
		chatRepository:                 chatRepository,
		channel:                        make(map[string][]chan *chat.Message),
	}
}

func (s *Server) TypingChannel(server pb.ChatService_TypingChannelServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetRoomId(ctx context.Context, req *pb.GetRoomIdReq) (*pb.RoomId, error) {
	//TODO implement me
	return &pb.RoomId{}, nil
}

func (s *Server) AddToken(ctx context.Context, req *pb.AddTokenReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.UserIdReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}
