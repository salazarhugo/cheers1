package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func GetUserId(customerId string) (string, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return "", err
	}

	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Documents(ctx)
	doc, err := snapshot.Next()
	if err != nil {
		return "", err
	}

	return doc.Ref.ID, nil
}

func GetOrder(orderId string) (*order.Order, error) {
	ctx := context.Background()

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	conn, err := grpc.DialContext(ctx, "order-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	order, err := client.GetOrder(ctx, &order.GetOrderRequest{OrderId: orderId})
	if err != nil {
		return nil, err
	}

	return order.GetOrder(), nil
}

func CalculateTotalPrice(tickets []*ticketpb.Ticket) (int64, error) {
	var amount int64 = 0

	for _, ticket := range tickets {
		amount += ticket.Price
	}

	return amount, nil
}
