package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (o ticketRepository) CreateTicket(
	userID string,
	ticket *ticketpb.Ticket,
) (*ticketpb.Ticket, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}

	doc := client.Collection("tickets").NewDoc()
	ticket.Id = doc.ID

	m, _ := utils.ProtoToMap(ticket)
	_, err = doc.Set(ctx, m)
	_, err = client.Collection("users").Doc(ticket.UserId).Collection("tickets").Doc(doc.ID).Set(ctx, m)

	if err != nil {
		return nil, err
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("ticket-topic", &ticketpb.TicketEvent{
			Event: &ticketpb.TicketEvent_Create{
				Create: &ticketpb.CreateTicket{
					Ticket: ticket,
					User:   nil,
				},
			},
		})
		if err != nil {
		}
	}()

	return ticket, nil
}
