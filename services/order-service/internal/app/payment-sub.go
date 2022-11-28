package app

import (
	"encoding/json"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/services/order-service/internal/repository"
	"github.com/stripe/stripe-go/v72/paymentintent"
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

	log.Println(event)

	pi, err := paymentintent.Get(
		event.PaymentIntentId,
		nil,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to get payment intent", http.StatusInternalServerError)
		return
	}
	repo := repository.NewOrderRepository()
	err = repo.UpdateOrder(&pb.Order{
		Status: string(pi.Status),
		Id:     pi.ID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}
	return
}
