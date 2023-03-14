package repository

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhook"
	"io"
	"net/http"
	"os"
)

func ValidateStripeRequest(
	w http.ResponseWriter,
	request *http.Request,
) *stripe.Event {
	const MaxBodyBytes = int64(65536)
	request.Body = http.MaxBytesReader(w, request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		msg := fmt.Sprintf("Error reading request body: %v\n", err)
		http.Error(w, msg, http.StatusServiceUnavailable)
		return nil
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		msg := fmt.Sprintf("⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return nil
	}

	endpointSecret := os.Getenv("STRIPE_WH")
	signatureHeader := request.Header.Get("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)

	if err != nil {
		msg := fmt.Sprintf("⚠️  Webhook signature verification failed. %v\n", err)
		http.Error(w, msg, http.StatusBadRequest)
		return nil
	}

	return &event
}
