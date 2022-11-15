package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"sync"
)

type Server struct {
	notification.UnimplementedNotificationServiceServer
	mu sync.Mutex
}

func NewServer() *Server {
	return &Server{
		notification.UnimplementedNotificationServiceServer{},
		sync.Mutex{},
	}
}
