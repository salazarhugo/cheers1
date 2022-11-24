package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateOrder(
	ctx context.Context,
	request *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	OrderReq := request.GetOrder()
	log.Println(OrderReq)
	if OrderReq == nil {
		return nil, status.Error(codes.InvalidArgument, "Order parameter can't be nil")
	}

	response, err := s.orderRepository.CreateOrder(userID, OrderReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create Order")
	}

	return &pb.CreateOrderResponse{
		Order: response,
	}, nil
}

func ValidateCreateOrderRequest(request *pb.CreateOrderRequest) error {
	order := request.GetOrder()
	if order == nil {
		return status.Error(codes.InvalidArgument, "missing parameter order")
	}
	return nil
}
