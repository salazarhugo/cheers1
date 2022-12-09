package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListTicket(
	ctx context.Context,
	request *pb.ListTicketRequest,
) (*pb.ListTicketResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	err = ValidateListTicketRequest(request)
	if err != nil {
		return nil, err
	}

	ticketList, err := s.ticketRepository.ListTicket(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.ListTicketResponse{
		Tickets: ticketList,
	}, nil
}

func ValidateListTicketRequest(request *pb.ListTicketRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "request can't be nil")
	}
	return nil
}
