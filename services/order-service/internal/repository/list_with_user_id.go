package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/iterator"
)

func (o orderRepository) ListOrderWithUserId(
	userID string,
) ([]*pb.Order, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	docs := client.Collection("orders").Where("partyHostId", "==", userID).Documents(ctx)

	orderList := make([]*pb.Order, 0)

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		order := &pb.Order{}
		err = utils.MapToProto(order, doc.Data())
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}

	return orderList, nil
}
