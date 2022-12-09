package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreatePayment(
	ctx context.Context,
	request *pb.CreatePaymentRequest,
) (*pb.CreatePaymentResponse, error) {
	userId, err := utils.GetUserId(ctx)
	var customerID *string = nil

	// If user is logged in, get the associated customer
	if err != nil {
		customer, _ := repository.GetCustomer(userId)
		if customer != nil {
			customerID = &customer.Id
		}
	}

	// Validate payment intent request
	err = ValidateCreatePaymentRequest(request)
	if err != nil {
		return nil, err
	}

	// Get tickets
	tickets, err := repository.GetTickets(
		request.GetTickets(),
		request.GetPartyId(),
	)

	// Calculate the total amount
	amount, err := repository.CalculateAmount(tickets)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		Customer: customerID,
	}

	// Create stripe payment intent
	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	party, err := repository.GetParty(request.PartyId)
	if err != nil {
		return nil, err
	}

	// Store the order in firestore
	err = repository.CreateOrder(
		paymentIntent,
		request,
		userId,
		party,
		tickets,
	)
	if err != nil {
		return nil, err
	}

	// Store payment intent
	return &pb.CreatePaymentResponse{
		ClientSecret: paymentIntent.ClientSecret,
	}, nil
}

func ValidateCreatePaymentRequest(
	request *pb.CreatePaymentRequest,
) error {
	if request.GetFirstName() == "" {
		return status.Error(codes.InvalidArgument, "missing parameter first_name")
	}
	if request.LastName == "" {
		return status.Error(codes.InvalidArgument, "missing parameter last_name")
	}
	if request.Email == "" {
		return status.Error(codes.InvalidArgument, "missing parameter email")
	}
	if request.PartyId == "" {
		return status.Error(codes.InvalidArgument, "missing parameter party_id")
	}
	if len(request.Tickets) < 1 {
		return status.Error(codes.InvalidArgument, "empty parameter tickets")
	}
	return nil
}
