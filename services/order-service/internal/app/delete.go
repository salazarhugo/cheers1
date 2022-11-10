package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
)

func (s *Server) DeleteOrder(
	ctx context.Context,
	request *pb.DeleteOrderRequest,
) (*pb.DeleteOrderResponse, error) {
	return &pb.DeleteOrderResponse{}, nil
}
