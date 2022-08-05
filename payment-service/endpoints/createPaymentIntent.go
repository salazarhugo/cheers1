package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"net/http"
	"os"
	"salazar/cheers/payment/utils"
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

	userDoc, err := client.Collection("stripe_customers").Doc(userId).Get(ctx)
	if err != nil {
	}

	var customer StripeCustomer
	if err := userDoc.DataTo(&customer); err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000),
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"clientSecret": paymentIntent.ClientSecret,
	})
}
