package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"log"
	"net/http"
	"os"
	"salazar/cheers/payment/utils"
)

type StripeCustomer struct {
	Id string `firestore:"customer_id"`
}

type PaymentIntentReq struct {
	Amount int64 `json:"amount" structs:"amount"`
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
	var amount int64 = 0

	paymentIntentReq := make(map[string]int64)

	err = json.NewDecoder(cc.Request().Body).Decode(&paymentIntentReq)
	if err != nil {
		return err
	}

	log.Println(paymentIntentReq)
	for ticketId, quantity := range paymentIntentReq {
		log.Println(ticketId, quantity)
		docsnap, err := client.Collection("ticketing").Doc(partyId).Collection("tickets").Doc(ticketId).Get(ctx)
		if err != nil {
			log.Println(err)
			return err
		}
		ticketMap := docsnap.Data()
		amount += ticketMap["price"].(int64) * quantity
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

	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	ref := client.Collection("stripe_customers").Doc(customer.Id).Collection("paymentIntents")
	_, _, err = ref.Add(ctx, map[string]interface{}{
		"id":         paymentIntent.ID,
		"amount":     paymentIntent.Amount,
		"created":    paymentIntent.Created,
		"customerId": paymentIntent.Customer.ID,
		"status":     paymentIntent.Status,
		"partyId":    partyId,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"clientSecret": paymentIntent.ClientSecret,
	})
}
