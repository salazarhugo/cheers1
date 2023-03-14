package events

import (
	"encoding/json"
	"fmt"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/stripe/stripe-go/v72"
	"net/http"
	"os"
)

func HandleStripeEvent(
	w http.ResponseWriter,
	r *http.Request,
) {
	session := utils.GetSession(utils.GetDriver())
	defer session.Close()

	stripe.Key = os.Getenv("STRIPE_SK")

	event := repository.ValidateStripeRequest(w, r)

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := parseEvent(event, &paymentIntent)
		if err != nil {
			http.Error(w, "no content", http.StatusInternalServerError)
			return
		}
		repository.HandlePaymentSuccess(paymentIntent)

	case "charge.refunded":
		var charge stripe.Charge
		err := parseEvent(event, &charge)
		if err != nil {
			http.Error(w, "no content", http.StatusInternalServerError)
			return
		}
		repository.HandleChargeRefunded(charge)

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := parseEvent(event, &paymentMethod)
		if err != nil {
			http.Error(w, "no content", http.StatusInternalServerError)
			return
		}
	default:
		msg := fmt.Sprintf("Unhandled event type: %s\n", event.Type)
		http.Error(w, msg, http.StatusNotImplemented)
		return
	}

	return
}

func parseEvent(event *stripe.Event, v any) error {
	err := json.Unmarshal(event.Data.Raw, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		return err
	}
	return nil
}
