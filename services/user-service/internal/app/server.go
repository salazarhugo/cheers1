package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"sync"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu             sync.Mutex
	userRepository repository.UserRepository
}

func NewServer() *Server {
	return &Server{
		pb.UnimplementedUserServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewUserRepository(utils.GetDriver()),
	}
}

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, "Failed retrieving metadata")
	}
	return md["user-id"][0], nil
}
