package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

func (o orderRepository) ListOrderWithPartyId(
	partyID string,
) ([]*pb.Order, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	docs := client.Collection("orders").Where("partyId", "==", partyID).Documents(ctx)

	orderList := make([]*pb.Order, 0)

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		order := &pb.Order{}
		bytes, err := json.Marshal(doc.Data())
		if err != nil {
			return nil, err
		}
		err = protojson.Unmarshal(bytes, order)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}

	return orderList, nil
}
