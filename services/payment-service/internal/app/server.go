package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"sync"
)

type Server struct {
	payment.UnimplementedPaymentServiceServer
	mu                sync.Mutex
	paymentRepository repository.PaymentRepository
}

func NewServer() *Server {
	return &Server{
		payment.UnimplementedPaymentServiceServer{},
		sync.Mutex{},
		repository.NewPaymentRepository(),
	}
}
