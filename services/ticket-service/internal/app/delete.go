package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
)

func (s *Server) DeleteTicket(
	ctx context.Context,
	request *pb.DeleteTicketRequest,
) (*pb.DeleteTicketResponse, error) {
	return &pb.DeleteTicketResponse{}, nil
}
