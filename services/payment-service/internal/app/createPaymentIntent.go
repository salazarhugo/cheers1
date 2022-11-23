package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"os"
	"time"
)

type StripeCustomer struct {
	Id string `firestore:"customer_id"`
}

func CreatePaymentIntent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	secret := os.Getenv("STRIPE_SK")
	stripe.Key = secret

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userId := cc.Get("userId").(string)
	partyId := cc.QueryParam("partyId")

	paymentIntentReq := make(map[string]int64)

	err = json.NewDecoder(cc.Request().Body).Decode(&paymentIntentReq)
	if err != nil {
		return err
	}

	log.Println(paymentIntentReq)
	tickets, err := getTickets(client, paymentIntentReq, partyId)
	log.Println(tickets)
	amount, err := calculateTotalPrice(tickets)
	log.Println(amount)
	if err != nil {
		log.Println(err)
		return err
	}

	userDoc, err := client.Collection("stripe_customers").Doc(userId).Get(ctx)
	if err != nil {
		return err
	}

	var customer StripeCustomer
	if err := userDoc.DataTo(&customer); err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	party, err := repository.GetParty(partyId)
	if err != nil {
		return err
	}

	// Store payment intent
	ref := client.Collection("orders").Doc(paymentIntent.ID)
	order := &order.Order{
		Id:          paymentIntent.ID,
		Status:      "",
		Amount:      paymentIntent.Amount,
		CustomerId:  paymentIntent.Customer.ID,
		UserId:      "",
		PartyId:     partyId,
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"clientSecret": paymentIntent.ClientSecret,
	})
}

func getTickets(
	client *firestore.Client,
	paymentIntentReq map[string]int64,
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
		var i int64
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
