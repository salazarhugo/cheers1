package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func GetUsers(userIds []string) ([]*pb.UserItem, error) {
	ctx := context.Background()
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	conn, err := grpc.DialContext(ctx, "user-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	response, err := client.GetUserItemsIn(ctx, &user.GetUserItemsInRequest{UserIds: userIds})
	if err != nil {
		log.Println(err)
	}
	return response.GetUsers(), nil
}

func GetUser(userId string) (*pb.UserItem, error) {
	response, err := GetUsers([]string{userId})
	if err != nil || len(response) < 1 {
		log.Println(err)
		return nil, err
	}

	return response[0], nil
}
