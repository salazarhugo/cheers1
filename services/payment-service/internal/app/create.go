package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
)

type StripeCustomer struct {
	Id string `firestore:"customer_id"`
}

func (s *Server) CreatePayment(
	ctx context.Context,
	request *pb.CreatePaymentRequest,
) (*pb.CreatePaymentResponse, error) {
	userId, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	secret := os.Getenv("STRIPE_SK")
	stripe.Key = secret

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}

	tickets, err := getTickets(client, request.GetTickets(), request.GetPartyId())
	log.Println(tickets)
	amount, err := calculateTotalPrice(tickets)
	log.Println(amount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userDoc, err := client.Collection("stripe_customers").Doc(userId).Get(ctx)
	if err != nil {
		return nil, err
	}

	var customer StripeCustomer
	if err := userDoc.DataTo(&customer); err != nil {
		return nil, status.Error(codes.NotFound, "customer not found")
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		Customer: &customer.Id,
	}

	// Create stripe payment intent
	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	party, err := repository.GetParty(request.GetPartyId())
	if err != nil {
		return nil, err
	}

	// Store payment intent
	ref := client.Collection("orders").Doc(paymentIntent.ID)
	order := &order.Order{
		Id:          paymentIntent.ID,
		Amount:      paymentIntent.Amount,
		CustomerId:  paymentIntent.Customer.ID,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PartyId:     request.PartyId,
		PartyHostId: party.HostId,
		CreateTime:  timestamppb.New(time.Unix(paymentIntent.Created, 0)),
		Tickets:     tickets,
	}

	m, err := utils.ProtoToMap(order)
	if err != nil {
		log.Println(err)
	}

	_, err = ref.Set(ctx, m)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreatePaymentResponse{
		ClientSecret: paymentIntent.ClientSecret,
	}, nil
}

func getTickets(
	client *firestore.Client,
	paymentIntentReq map[string]uint32,
	partyId string,
) ([]*ticketpb.Ticket, error) {
	tickets := make([]*ticketpb.Ticket, 0, 0)

	for ticketId, quantity := range paymentIntentReq {
		doc, err := client.Collection("ticketing").Doc(partyId).Collection("tickets").Doc(ticketId).Get(context.Background())
		if err != nil {
			return nil, err
		}
		ticket := &ticketpb.Ticket{}
		err = utils.MapToProto(ticket, doc.Data())
		if err != nil {
			return nil, err
		}
		var i uint32
		for i = 0; i < quantity; i++ {
			tickets = append(tickets, ticket)
		}
	}

	return tickets, nil
}

func calculateTotalPrice(
	tickets []*ticketpb.Ticket,
) (int64, error) {
	var amount int64 = 0

	for _, ticket := range tickets {
		amount += ticket.Price
	}

	return amount, nil
}
