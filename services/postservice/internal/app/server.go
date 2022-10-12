package app

import (
	"context"
	"github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/postservice/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"sync"
)

type Server struct {
	post.UnimplementedPostServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu              sync.Mutex
	partyRepository repository.PostRepository
}

func NewServer() *Server {
	return &Server{
		post.UnimplementedPostServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewPostRepository(utils.GetDriver()),
	}
}

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, "Failed retrieving metadata")
	}
	return md["user-id"][0], nil
}
