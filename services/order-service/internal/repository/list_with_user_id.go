package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	doc, err := client.Collection("stripe_customers").Doc(userID).Get(ctx)

	if doc.Exists() == false {
		return nil, status.Error(codes.Internal, "user is not a stripe customer")
	}

	customerId := doc.Data()["customer_id"]
	docs := client.Collection("orders").Where("customerId", "==", customerId).Documents(ctx)

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
