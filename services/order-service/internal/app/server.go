package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/order-service/internal/repository"
	"sync"
)

type Server struct {
	order.UnimplementedOrderServiceServer
	mu              sync.Mutex
	orderRepository repository.OrderRepository
}

func NewServer() *Server {
	app := utils.InitializeAppDefault()
	client, _ := app.Firestore(context.Background())
	return &Server{
		order.UnimplementedOrderServiceServer{},
		sync.Mutex{},
		repository.NewOrderRepository(client),
	}
}
