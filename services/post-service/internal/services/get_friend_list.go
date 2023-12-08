package services

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func GetFriendList(userId string) ([]string, error) {
	ctx := context.Background()

	conn := utils.CreateServiceConnection(ctx, "friendship-service-r3a2dr4u4a-nw.a.run.app:443")
	defer conn.Close()

	client := friendship.NewFriendshipServiceClient(conn)

	response, err := client.ListFriendIds(ctx, &friendship.ListFriendIdsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	return response.GetIds(), nil
}
