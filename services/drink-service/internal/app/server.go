package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/services/drink-service/internal/repository"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	drink.UnimplementedDrinkServiceServer
	mu         sync.Mutex
	repository repository.Repository
	logger     *logrus.Logger
}

func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		drink.UnimplementedDrinkServiceServer{},
		sync.Mutex{},
		repository.NewRepository(),
		logger,
	}
}
