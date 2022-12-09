package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (o ticketRepository) GetTicket(id string) (*pb.Ticket, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("tickets").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	ticket := &pb.Ticket{}
	err = utils.MapToProto(ticket, doc.Data())
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
