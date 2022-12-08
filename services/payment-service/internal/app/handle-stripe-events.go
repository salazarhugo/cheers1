package app

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/stripe/stripe-go/v72"
	"net/http"
	"os"
)

func HandleStripeEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	stripe.Key = os.Getenv("STRIPE_SK")

	event, err := repository.ValidateStripeRequest(cc)
	if err != nil {
		return err
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := parseEvent(event, &paymentIntent)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		repository.HandlePaymentSuccess(paymentIntent)

	case "charge.refunded":
		var charge stripe.Charge
		err := parseEvent(event, &charge)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		repository.HandleChargeRefunded(charge)

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := parseEvent(event, &paymentMethod)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		return cc.NoContent(http.StatusNotImplemented)
	}

	return cc.NoContent(http.StatusOK)
}

func parseEvent(event *stripe.Event, v any) error {
	err := json.Unmarshal(event.Data.Raw, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		return err
	}
	return nil
}
