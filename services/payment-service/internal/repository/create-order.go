package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	partypb "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/stripe/stripe-go/v72"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func CreateOrder(
	paymentIntent *stripe.PaymentIntent,
	request *payment.CreatePaymentRequest,
	userID string,
	party *partypb.Party,
	tickets []*ticketpb.Ticket,
) error {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}

	ref := client.Collection("orders").Doc(paymentIntent.ID)
	paymentMethodType := ""
	if paymentIntent.PaymentMethod != nil {
		paymentMethodType = string(paymentIntent.PaymentMethod.Type)
	}

	customerId := ""
	if paymentIntent.Customer != nil {
		customerId = paymentIntent.Customer.ID
	}

	order := &order.Order{
		Id:                 paymentIntent.ID,
		PaymentMethodTypes: paymentIntent.PaymentMethodTypes,
		PaymentMethodType:  paymentMethodType,
		Status:             string(paymentIntent.Status),
		Amount:             paymentIntent.Amount,
		CustomerId:         customerId,
		UserId:             userID,
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		Email:              request.Email,
		PartyId:            request.PartyId,
		PartyName:          party.Name,
		PartyHostId:        party.HostId,
		CreateTime:         paymentIntent.Created,
		Tickets:            tickets,
	}

	m, err := utils.ProtoToMap(order)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ref.Set(ctx, m)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
