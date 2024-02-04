package service

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func GetUser(userId string) (*user.User, error) {
	ctx := context.Background()

	conn := utils.CreateServiceConnection(ctx, "user-service-r3a2dr4u4a-nw.a.run.app:443")
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	response, err := client.GetUserNode(ctx, &userpb.GetUserNodeRequest{UserId: userId})
	if err != nil {
		log.Println(err)
	}
	return response.GetUser(), nil
}
