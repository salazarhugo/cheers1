package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListOrganizerOrders(
	ctx context.Context,
	request *pb.ListOrganizerOrdersRequest,
) (*pb.ListOrganizerOrdersResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	orderList, err := s.orderRepository.ListOrganizationOrders(userID, request.Query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.ListOrganizerOrdersResponse{
		Orders: orderList,
	}, nil
}
