package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"sync"
)

type Server struct {
	notification.UnimplementedNotificationServiceServer
	mu         sync.Mutex
	repository repository.Repository
}

func NewServer() *Server {
	return &Server{
		notification.UnimplementedNotificationServiceServer{},
		sync.Mutex{},
		repository.NewRepository(),
	}
}
