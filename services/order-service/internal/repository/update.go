package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (o orderRepository) UpdateOrder(
	order *pb.Order,
) error {
	log.Printf("Updating order with status %s", order.Status)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}
	defer client.Close()

	m, err := utils.ProtoToMap(order)
	if err != nil {
		return err
	}
	_, err = client.Collection("orders").Doc(order.Id).Set(ctx, m, firestore.MergeAll)

	return err
}
