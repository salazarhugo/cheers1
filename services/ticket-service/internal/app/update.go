package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTicket(
	ctx context.Context,
	request *pb.UpdateTicketRequest,
) (*pb.UpdateTicketResponse, error) {
	err := ValidateUpdateTicketRequest(request)
	if err != nil {
		return nil, err
	}

	TicketReq := request.GetTicket()

	err = s.ticketRepository.UpdateTicket(TicketReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create Ticket")
	}

	Ticket, err := s.ticketRepository.GetTicket(TicketReq.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get Ticket")
	}

	return &pb.UpdateTicketResponse{
		Ticket: Ticket,
	}, nil
}

func ValidateUpdateTicketRequest(request *pb.UpdateTicketRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "Ticket parameter can't be nil")
	}

	Ticket := request.GetTicket()

	if Ticket.Id == "" {
		return status.Error(codes.InvalidArgument, "missing Ticket id")
	}

	return nil
}
