package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/services/drink-service/internal/repository"
	"sync"
)

type Server struct {
	drink.UnimplementedDrinkServiceServer
	mu         sync.Mutex
	repository repository.Repository
}

func NewServer() *Server {
	return &Server{
		drink.UnimplementedDrinkServiceServer{},
		sync.Mutex{},
		repository.NewRepository(),
	}
}
