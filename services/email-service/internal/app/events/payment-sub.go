package events

import (
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/email-service/internal/emails"
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
	"log"
	"net/http"
)

func PaymentSub(w http.ResponseWriter, r *http.Request) {
	event := &payment.PaymentEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	order := event.Order
	totalPrice, err := repository.CalculateTotalPrice(order.Tickets)

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to calculate amount", http.StatusInternalServerError)
		return
	}

	tickets := make([]*emails.EmailTicket, 0)

	for i := 0; i < len(order.Tickets); i++ {
		ticket := emails.New(
			order.Tickets[i].Name,
			order.Tickets[i].Quantity,
			order.Tickets[i].Price,
		)
		tickets = append(tickets, ticket)
	}

	err = emails.SendPartyReceiptEmail(
		order.Email,
		order.FirstName,
		order.PartyName,
		tickets,
	)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(totalPrice)
}
