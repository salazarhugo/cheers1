package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func GetTickets(
	paymentIntentReq map[string]uint32,
	partyId string,
) ([]*ticketpb.Ticket, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	tickets := make([]*ticketpb.Ticket, 0, 0)

	for ticketId, quantity := range paymentIntentReq {
		doc, err := client.Collection("ticketing").Doc(partyId).Collection("tickets").Doc(ticketId).Get(context.Background())
		if err != nil {
			return nil, err
		}
		ticket := &ticketpb.Ticket{}
		err = mapper.MapToProto(ticket, doc.Data())
		if err != nil {
			return nil, err
		}
		var i uint32
		for i = 0; i < quantity; i++ {
			tickets = append(tickets, ticket)
		}
	}

	return tickets, nil
}
