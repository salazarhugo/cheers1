package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (o orderRepository) GetOrder(id string) (*pb.Order, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("orders").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	order := &pb.Order{}
	err = utils.MapToProto(order, doc.Data())
	if err != nil {
		return nil, err
	}

	return order, nil
}
