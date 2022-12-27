package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/services/comment-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

type Server struct {
	comment.UnimplementedCommentServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu                sync.Mutex
	commentRepository repository.Repository
}

func NewServer() *Server {
	return &Server{
		comment.UnimplementedCommentServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewRepository(),
	}
}
