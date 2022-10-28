package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	user "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

func (r repository) SendChatNotification(authorization string) error {
	context := context.Background()

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return err
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	conn, err := grpc.DialContext(context, "android-gateway-clzdlli7.nw.gateway.dev:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	ctx := metadata.AppendToOutgoingContext(context, "authorization", authorization)

	response, err := client.GetUserItemsIn(ctx, &user.GetUserItemsInRequest{UserIds: []string{""}})
	if err != nil {
		log.Println(err)
	}

	log.Println(response)

	return nil
}
