package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
)

type TicketRepository interface {
	CreateTicket(userID string, pb *pb.Ticket) (*pb.Ticket, error)
	GetTicket(id string) (*pb.Ticket, error)
	UpdateTicket(pb *pb.Ticket) error
	DeleteTicket(userID string, orderID string) error

	ListTicket(userID string) ([]*pb.Ticket, error)
}

type ticketRepository struct {
}

func NewTicketRepository() TicketRepository {
	return &ticketRepository{}
}
