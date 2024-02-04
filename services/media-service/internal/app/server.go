package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/media/v1"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

type Server struct {
	media.UnimplementedMediaServiceServer
	grpc_health_v1.UnimplementedHealthServer
	mu              sync.Mutex
	mediaRepository repository.MediaRepository
}

func NewServer() *Server {
	return &Server{
		media.UnimplementedMediaServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		sync.Mutex{},
		repository.NewPostRepository(),
	}
}
