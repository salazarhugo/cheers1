package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (o ticketRepository) UpdateTicket(
	ticket *ticketpb.Ticket,
) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}
	defer client.Close()

	m, err := utils.ProtoToMap(ticket)
	if err != nil {
		return err
	}
	_, err = client.Collection("tickets").Doc(ticket.Id).Set(ctx, m, firestore.MergeAll)
	if err != nil {
		return err
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("ticket-topic", &ticketpb.TicketEvent{
			Event: &ticketpb.TicketEvent_Update{
				Update: &ticketpb.UpdateTicket{
					Ticket: ticket,
					User:   nil,
				},
			},
		})
		if err != nil {
		}
	}()
	return err
}
