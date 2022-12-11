package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"log"
)

func (s *Server) ListPartyOrders(
	ctx context.Context,
	request *pb.ListPartyOrdersRequest,
) (*pb.ListPartyOrdersResponse, error) {
	orderList, err := s.orderRepository.ListPartyOrders(request.PartyId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.ListPartyOrdersResponse{
		Orders: orderList,
	}, nil
}
