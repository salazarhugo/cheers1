package repository

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c chatRepository) DeleteRoom(
	userID string,
	roomID string,
) error {
	room, err := c.cache.GetRoomWithId(userID, roomID)
	if err != nil {
		return err
	}

	var isAdmin bool = false
	for _, admin := range room.Admins {
		if admin == userID {
			isAdmin = true
			break
		}
	}

	if isAdmin != true {
		return status.Error(codes.PermissionDenied, "you are not group admin")
	}

	return c.cache.DeleteRoom(roomID)
}
