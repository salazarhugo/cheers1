package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/services/payment-service/utils"
	"google.golang.org/api/iterator"
)

func (o orderRepository) ListOrderWithUserId(
	userID string,
) ([]*pb.Order, error) {
	ctx := context.Background()

	docs := o.client.Collection("orders").Where("userId", "==", userID).OrderBy("createTime", firestore.Desc).Documents(ctx)

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
