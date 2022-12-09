package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/services/ticket-service/internal/repository"
	"sync"
)

type Server struct {
	ticket.UnimplementedTicketServiceServer
	mu               sync.Mutex
	ticketRepository repository.TicketRepository
}

func (s *Server) mustEmbedUnimplementedTicketServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewServer() *Server {
	return &Server{
		ticket.UnimplementedTicketServiceServer{},
		sync.Mutex{},
		repository.NewTicketRepository(),
	}
}
