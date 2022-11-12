package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/salazarhugo/cheers1/services/payment-service/utils"
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

	// Store payment intent
	ref := client.Collection("orders").Doc(paymentIntent.ID)
	_, err = ref.Set(ctx, map[string]interface{}{
		"id":         paymentIntent.ID,
		"amount":     paymentIntent.Amount,
		"customerId": paymentIntent.Customer.ID,
		"status":     paymentIntent.Status,
		"partyId":    partyId,
		"tickets":    tickets,
		"createTime": timestamppb.New(time.Unix(paymentIntent.Created, 0)),
	})

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
) ([]map[string]interface{}, error) {
	tickets := make([]map[string]interface{}, 0, 0)

	for ticketId, quantity := range paymentIntentReq {
		docsnap, err := client.Collection("ticketing").Doc(partyId).Collection("tickets").Doc(ticketId).Get(context.Background())
		if err != nil {
			return nil, err
		}
		ticketMap := docsnap.Data()
		var i int64
		for i = 0; i < quantity; i++ {
			tickets = append(tickets, ticketMap)
		}
	}

	return tickets, nil
}

func calculateTotalPrice(
	tickets []map[string]interface{},
) (int64, error) {
	var amount int64 = 0

	for _, ticket := range tickets {
		amount += ticket["price"].(int64)
	}

	return amount, nil
}
