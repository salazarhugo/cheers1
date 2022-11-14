package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/services/account-service/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "order-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()
	client := order.NewOrderServiceClient(conn)
	response, err := client.GetOrder(ctx, &order.GetOrderRequest{OrderId: event.GetPaymentIntentId()})
	if err != nil {
		log.Println(err)
		return
	}
	order := response.GetOrder()
	log.Println(order)
	partyId := order.GetPartyId()

	conn2, err := grpc.DialContext(ctx, "party-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()
	clientParty := party.NewPartyServiceClient(conn2)
	res, err := clientParty.GetParty(ctx, &party.GetPartyRequest{PartyId: partyId})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.GetParty())

	datastore, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		log.Println(err)
		return
	}
	defer datastore.Close()

	repository := repository.NewAccountRepository()
	err = repository.IncrementBalance(res.GetParty().HostId, order.Amount)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
