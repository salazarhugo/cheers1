package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/iterator"
)

func (o orderRepository) ListOrderWithPartyId(
	partyID string,
) ([]*pb.Order, error) {
	ctx := context.Background()

	docs := o.client.Collection("orders").Where("partyId", "==", partyID).OrderBy("createTime", firestore.Desc).Documents(ctx)

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
