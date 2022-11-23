package app

import (
	"context"
	"fmt"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListOrder(
	ctx context.Context,
	request *pb.ListOrderRequest,
) (*pb.ListOrderResponse, error) {
	err := ValidateListOrderRequest(request)
	if err != nil {
		return nil, err
	}

	switch filter := request.GetFilter().(type) {
	case *pb.ListOrderRequest_PartyId:
		orderList, err := s.orderRepository.ListOrderWithPartyId(filter.PartyId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &pb.ListOrderResponse{
			Orders: orderList,
		}, nil
	case *pb.ListOrderRequest_UserId:
		orderList, err := s.orderRepository.ListOrderWithUserId(filter.UserId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &pb.ListOrderResponse{
			Orders: orderList,
		}, nil
	default:
		fmt.Println("No matching operations")
	}

	return nil, err
}

func ValidateListOrderRequest(request *pb.ListOrderRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "request can't be nil")
	}
	return nil
}
