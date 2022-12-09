package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (o ticketRepository) UpdateTicket(
	ticket *pb.Ticket,
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

	return err
}
