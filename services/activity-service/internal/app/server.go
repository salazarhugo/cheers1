package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	activity.UnimplementedActivityServiceServer
	grpc_health_v1.UnimplementedHealthServer
	activityRepository repository.ActivityRepository
}

func (s *Server) mustEmbedUnimplementedActivityServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewServer() *Server {
	return &Server{
		activity.UnimplementedActivityServiceServer{},
		grpc_health_v1.UnimplementedHealthServer{},
		repository.NewRepository(),
	}
}
