package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
)

type OrderRepository interface {
	CreateOrder(userID string, pb *pb.Order) (*pb.Order, error)
	GetOrder(id string) (*pb.Order, error)
	UpdateOrder(pb *pb.Order) (string, error)
	DeleteOrder(id string) error
}

type orderRepository struct {
	driver neo4j.Driver
}

func NewOrderRepository(driver neo4j.Driver) OrderRepository {
	return &orderRepository{
		driver: driver,
	}
}
