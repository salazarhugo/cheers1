package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateTicket(
	ctx context.Context,
	request *pb.CreateTicketRequest,
) (*pb.CreateTicketResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	TicketReq := request.GetTicket()
	log.Println(TicketReq)
	if TicketReq == nil {
		return nil, status.Error(codes.InvalidArgument, "Ticket parameter can't be nil")
	}

	response, err := s.ticketRepository.CreateTicket(userID, TicketReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create Ticket")
	}

	return &pb.CreateTicketResponse{
		Ticket: response,
	}, nil
}

func ValidateCreateTicketRequest(request *pb.CreateTicketRequest) error {
	ticket := request.GetTicket()
	if ticket == nil {
		return status.Error(codes.InvalidArgument, "missing parameter ticket")
	}
	return nil
}
