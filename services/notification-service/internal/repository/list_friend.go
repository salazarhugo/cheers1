package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// ListFriend /*
func ListFriend(userId string) ([]string, error) {
	ctx := context.Background()
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	conn, err := grpc.DialContext(ctx, "friendship-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := friendship.NewFriendshipServiceClient(conn)

	response, err := client.ListFriend(ctx, &friendship.ListFriendRequest{UserId: userId})
	if err != nil {
		log.Println(err)
	}

	friendUUIDs := make([]string, 0)

	for _, friend := range response.Items {
		friendUUIDs = append(friendUUIDs, friend.Id)
	}

	return friendUUIDs, nil
}
