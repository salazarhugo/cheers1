package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

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
		grpc.WithUnaryInterceptor(utils.CloudRunInterceptor),
	)
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	order, err := client.GetOrder(ctx, &order.GetOrderRequest{OrderId: orderId})
	if err != nil {
		return nil, err
	}

	return order.GetOrder(), nil
}
