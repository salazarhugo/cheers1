package repository

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (c chatRepository) ListMembers(
	context context.Context,
	request *pb.ListMembersRequest,
) ([]*user.UserItem, error) {
	membersIDs, err := c.cache.GetRoomMembers(request.RoomId)

	users, err := c.GetUsersNode(membersIDs)
	if err != nil {
		return nil, err
	}

	items := make([]*user.UserItem, 0)
	for _, u := range users {
		items = append(items, u.ToUserItemPb())
	}

	return items, nil
}
