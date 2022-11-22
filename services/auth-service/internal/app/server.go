package app

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu             sync.Mutex
	authRepository repository.AuthRepository
}

func NewServer() *Server {
	return &Server{
		pb.UnimplementedAuthServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewAuthRepository(),
	}
}
