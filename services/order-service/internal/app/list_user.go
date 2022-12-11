package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListUserOrders(
	ctx context.Context,
	request *pb.ListUserOrdersRequest,
) (*pb.ListUserOrdersResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	orderList, err := s.orderRepository.ListUserOrders(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.ListUserOrdersResponse{
		Orders: orderList,
	}, nil
}
