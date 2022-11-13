package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"github.com/salazarhugo/cheers1/services/account-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"sync"
)

type Server struct {
	pb.UnimplementedAccountServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu                sync.Mutex
	accountRepository repository.AccountRepository
}

func NewServer() *Server {
	return &Server{
		pb.UnimplementedAccountServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewAccountRepository(),
	}
}

func GetAccountId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, "Failed retrieving metadata")
	}
	return md["account-id"][0], nil
}
