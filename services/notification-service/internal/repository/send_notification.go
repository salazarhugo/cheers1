package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"firebase.google.com/go/v4/messaging"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func (r repository) SendChatNotification(userWithToken map[string][]string) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	fcmClient, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	for userId, tokens := range userWithToken {
		log.Println(userId)
		response, _ := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: "Post Like",
				Body:  "someone has liked your post",
			},
			Tokens: tokens,
		})
		log.Println("Response success count : ", response.SuccessCount)
		log.Println("Response failure count : ", response.FailureCount)
	}

	return nil
}

func GetUsers() ([]*pb.UserItem, error) {
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

	response, err := client.GetUserItemsIn(ctx, &user.GetUserItemsInRequest{UserIds: []string{""}})
	if err != nil {
		log.Println(err)
	}
	return response.GetUsers(), nil
}
