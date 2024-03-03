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

	items, err := c.cache.ListUser(membersIDs)
	if err != nil {
		return nil, err
	}

	return items, nil
}
