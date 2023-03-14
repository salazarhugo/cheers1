package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	friendshihpb "github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func ListFriends(userId string) ([]*user.UserItem, error) {
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

	client := friendshihpb.NewFriendshipServiceClient(conn)

	response, err := client.ListFriend(ctx, &friendshihpb.ListFriendRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
