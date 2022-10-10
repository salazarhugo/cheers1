package app

import (
	"context"
	"github.com/salazarhugo/cheers1/genproto/grpc/health/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Check is for health checking.
func (s *Server) Check(
	ctx context.Context,
	req *health.HealthCheckRequest,
) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(
	req *health.HealthCheckRequest,
	ws health.Health_WatchServer,
) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
