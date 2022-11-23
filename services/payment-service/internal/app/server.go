package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"sync"
)

type Server struct {
	payment.UnimplementedPaymentServiceServer
	mu                sync.Mutex
	paymentRepository repository.PaymentRepository
}

func NewServer() *Server {
	app := utils.InitializeAppDefault()
	client, _ := app.Firestore(context.Background())
	return &Server{
		payment.UnimplementedPaymentServiceServer{},
		sync.Mutex{},
		repository.NewPaymentRepository(client),
	}
}
