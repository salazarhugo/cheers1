package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/refund"
	"log"
)

func (s *Server) RefundPayment(
	ctx context.Context,
	request *pb.RefundPaymentRequest,
) (*pb.RefundPaymentResponse, error) {

	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(request.PaymentIntentId),
	}
	result, err := refund.New(params)
	log.Println(result)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.RefundPaymentResponse{}, nil
}
