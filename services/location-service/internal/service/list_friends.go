package service

import (
	"context"
	friendshihpb "github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func ListFriends(ctx context.Context, userId string) ([]*user.UserItem, error) {
	conn := utils.CreateServiceConnection(ctx, "friendship-service-r3a2dr4u4a-nw.a.run.app:443")
	defer conn.Close()

	client := friendshihpb.NewFriendshipServiceClient(conn)

	response, err := client.ListFriend(ctx, &friendshihpb.ListFriendRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
