package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

type Server struct {
	friendship.UnimplementedFriendshipServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu                   sync.Mutex
	friendshipRepository repository.Repository
}

func NewServer() *Server {
	return &Server{
		friendship.UnimplementedFriendshipServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewRepository(),
	}
}
