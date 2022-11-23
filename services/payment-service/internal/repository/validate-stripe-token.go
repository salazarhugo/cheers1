package repository

import (
	"encoding/json"
	"fmt"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhook"
	"io"
	"net/http"
	"os"
)

func ValidateStripeRequest(cc *utils.CustomContext) (*stripe.Event, error) {
	const MaxBodyBytes = int64(65536)
	cc.Request().Body = http.MaxBytesReader(cc.Response().Writer, cc.Request().Body, MaxBodyBytes)
	payload, err := io.ReadAll(cc.Request().Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		return nil, cc.NoContent(http.StatusServiceUnavailable)
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		return nil, cc.NoContent(http.StatusBadRequest)
	}

	endpointSecret := os.Getenv("STRIPE_WH")
	signatureHeader := cc.Request().Header.Get("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
		return nil, cc.NoContent(http.StatusBadRequest)
	}

	return &event, nil
}
