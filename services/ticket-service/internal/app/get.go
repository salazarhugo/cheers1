package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) GetTicket(
	ctx context.Context,
	request *pb.GetTicketRequest,
) (*pb.GetTicketResponse, error) {
	err := ValidateGetTicketRequest(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ticket, err := s.ticketRepository.GetTicket(request.TicketId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetTicketResponse{
		Ticket: ticket,
	}, nil
}

func ValidateGetTicketRequest(request *pb.GetTicketRequest) error {
	TicketId := request.GetTicketId()
	if TicketId == "" {
		return status.Error(codes.InvalidArgument, "Ticket_id parameter can't be nil")
	}

	return nil
}
