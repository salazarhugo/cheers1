package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/order-service/internal/repository"
	"sync"
)

type Server struct {
	order.UnimplementedOrderServiceServer
	mu              sync.Mutex
	orderRepository repository.OrderRepository
}

func NewServer() *Server {
	return &Server{
		order.UnimplementedOrderServiceServer{},
		sync.Mutex{},
		repository.NewOrderRepository(utils.GetDriver()),
	}
}
