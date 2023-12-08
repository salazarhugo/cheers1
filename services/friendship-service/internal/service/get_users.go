package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
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
		grpc.WithUnaryInterceptor(utils.CloudRunInterceptor),
	)
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	response, err := client.GetUsersIn(ctx, &user.GetUsersInRequest{UserIds: userIds})
	if err != nil {
		log.Println(err)
	}

	items := make([]*userpb.UserItem, 0)
	for _, user := range response.Users {
		items = append(items, &userpb.UserItem{
			Id:                 user.Id,
			Name:               user.Name,
			Username:           user.Username,
			Verified:           user.Verified,
			Picture:            user.Picture,
			HasFollowed:        false,
			StoryState:         0,
			Friend:             true,
			Requested:          false,
			HasRequestedViewer: false,
		})
	}
	return items, nil
}
