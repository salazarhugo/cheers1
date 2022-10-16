package app

import (
	"github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

type Server struct {
	post.UnimplementedPostServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu             sync.Mutex
	postRepository repository.PostRepository
}

func NewServer() *Server {
	return &Server{
		post.UnimplementedPostServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewPostRepository(utils.GetDriver()),
	}
}
