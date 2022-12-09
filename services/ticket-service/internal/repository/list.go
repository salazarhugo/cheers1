package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/iterator"
)

func (o ticketRepository) ListTicket(
	userID string,
) ([]*pb.Ticket, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	docs := client.Collection("users").Doc(userID).Collection("tickets").Documents(ctx)

	ticketList := make([]*pb.Ticket, 0)

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		ticket := &pb.Ticket{}
		err = utils.MapToProto(ticket, doc.Data())
		if err != nil {
			return nil, err
		}
		ticketList = append(ticketList, ticket)
	}

	return ticketList, nil
}
