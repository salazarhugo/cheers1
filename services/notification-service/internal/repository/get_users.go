package repository

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func GetUsers(userIds []string) ([]*pb.User, error) {
	ctx := context.Background()

	conn := utils.CreateServiceConnection(ctx, "user-service-r3a2dr4u4a-nw.a.run.app:443")
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	response, err := client.GetUsersIn(ctx, &user.GetUsersInRequest{UserIds: userIds})
	if err != nil {
		log.Println(err)
	}
	return response.GetUsers(), nil
}
