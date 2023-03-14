package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"github.com/salazarhugo/cheers1/services/location-service/internal/repository"
	"sync"
)

type Server struct {
	location.UnimplementedLocationServiceServer
	mu         sync.Mutex
	repository repository.Repository
}

func NewServer() *Server {
	return &Server{
		location.UnimplementedLocationServiceServer{},
		sync.Mutex{},
		repository.NewRepository(),
	}
}
