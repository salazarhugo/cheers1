package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/services/note-service/internal/domain"
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	note.UnimplementedNoteServiceServer
	mu         sync.Mutex
	repository repository.Repository
	domain     domain.Domain
	logger     *logrus.Logger
}

func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		note.UnimplementedNoteServiceServer{},
		sync.Mutex{},
		repository.NewRepository(),
		domain.NewDomain(),
		logger,
	}
}
