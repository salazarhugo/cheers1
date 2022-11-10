package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOrder(
	ctx context.Context,
	request *pb.UpdateOrderRequest,
) (*pb.UpdateOrderResponse, error) {
	err := ValidateUpdateOrderRequest(request)
	if err != nil {
		return nil, err
	}

	OrderReq := request.GetOrder()

	OrderID, err := s.orderRepository.UpdateOrder(OrderReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create Order")
	}

	Order, err := s.orderRepository.GetOrder(OrderID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get Order")
	}

	return &pb.UpdateOrderResponse{
		Order: Order,
	}, nil
}

func ValidateUpdateOrderRequest(request *pb.UpdateOrderRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "Order parameter can't be nil")
	}

	Order := request.GetOrder()

	if Order.Id == "" {
		return status.Error(codes.InvalidArgument, "missing Order id")
	}

	return nil
}
