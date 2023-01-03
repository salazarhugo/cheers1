package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func GetUserFollowers(userId string) ([]string, error) {
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

	response, err := client.ListFollowers(ctx, &user.ListFollowersRequest{UserId: userId})
	if err != nil {
		log.Println(err)
	}

	userIds := make([]string, 0)
	for _, user := range response.Users {
		userIds = append(userIds, user.Id)
	}

	return userIds, nil
}
