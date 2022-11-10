package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrder(
	ctx context.Context,
	request *pb.GetOrderRequest,
) (*pb.GetOrderResponse, error) {
	err := ValidateGetOrderRequest(request)
	if err != nil {
		return nil, err
	}

	Order, err := s.orderRepository.GetOrder(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Order: Order,
	}, nil
}

func ValidateGetOrderRequest(request *pb.GetOrderRequest) error {
	OrderId := request.GetOrderId()
	if OrderId == "" {
		return status.Error(codes.InvalidArgument, "Order_id parameter can't be nil")
	}

	return nil
}
