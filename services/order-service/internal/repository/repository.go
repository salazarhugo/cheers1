package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
)

type OrderRepository interface {
	CreateOrder(userID string, pb *pb.Order) (*pb.Order, error)
	GetOrder(id string) (*pb.Order, error)
	UpdateOrder(pb *pb.Order) error
	DeleteOrder(id string) error

	ListOrderWithPartyId(partyID string) ([]*pb.Order, error)
	ListOrderWithUserId(userID string) ([]*pb.Order, error)
}

type orderRepository struct {
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
