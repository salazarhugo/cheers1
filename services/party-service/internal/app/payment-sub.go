package app

import (
	"context"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"log"
	"net/http"
)

func PaymentSub(w http.ResponseWriter, r *http.Request) {
	event := &payment.PaymentEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		http.Error(w, "failed to parse pub/sub event", http.StatusInternalServerError)
	}
	log.Println(event)

	order := event.Order

	userID, err := GetUserId(order.CustomerId)
	partyID, err := GetPartyId(event.GetPaymentIntentId(), order.CustomerId)

	repo := repository.NewPartyRepository(utils.GetDriver())
	err = repo.GoingParty(userID, partyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func GetPartyId(paymentIntentId string, customerId string) (string, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return "", err
	}

	doc, err := client.Collection("orders").Doc(paymentIntentId).Get(ctx)
	if err != nil {
		return "", err
	}
	orderMap := doc.Data()
	partyId := orderMap["partyId"].(string)

	return partyId, nil
}

func GetUserId(customerId string) (string, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return "", err
	}

	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Documents(ctx)
	doc, err := snapshot.Next()
	if err != nil {
		return "", err
	}

	return doc.Ref.ID, nil
}
