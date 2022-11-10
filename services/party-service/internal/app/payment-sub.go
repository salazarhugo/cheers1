package app

import (
	"context"
	"encoding/json"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func PaymentSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad HTTP Request", http.StatusBadRequest)
		return
	}
	// byte slice unmarshalling handles base64 decoding.
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad HTTP Request", http.StatusBadRequest)
		return
	}

	event := &payment.PaymentEvent{}
	err = proto.Unmarshal(m.Message.Data, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	userID, err := GetUserId(event.GetCustomerId())
	partyID, err := GetPartyId(event.GetPaymentIntentId(), event.GetCustomerId())

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
