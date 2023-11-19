package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"log"
	"net/http"
)

func TicketSub(w http.ResponseWriter, r *http.Request) {
	event := &ticket.TicketEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		http.Error(w, "failed to parse pub/sub event", http.StatusInternalServerError)
	}
	log.Println(event)

	repo := repository.NewRepository()

	switch e := event.Event.(type) {
	case *ticket.TicketEvent_Create:
		ticket := e.Create.Ticket
		err = repo.UpdateMinimumPrice(ticket.Price, ticket.PartyId)
	case *ticket.TicketEvent_Delete:
		ticket := e.Delete.Ticket
		err = repo.UpdateMinimumPrice(-1, ticket.PartyId)
	case *ticket.TicketEvent_Update:
		ticket := e.Update.Ticket
		err = repo.UpdateMinimumPrice(ticket.Price, ticket.PartyId)
	default:
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
