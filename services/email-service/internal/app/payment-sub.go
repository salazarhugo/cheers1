package app

import (
	"encoding/json"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// byte slice unmarshalling handles base64 decoding.
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	event := &payment.PaymentEvent{}
	err = proto.Unmarshal(m.Message.Data, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	order, err := repository.GetOrder(event.GetPaymentIntentId())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get user tickets", http.StatusInternalServerError)
		return
	}

	totalPrice, _ := repository.CalculateTotalPrice(order.Tickets)

	SendEmail(event.Email, event.FirstName, order.Tickets, totalPrice)
}
