package repository

import (
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c chatRepository) DeleteRoom(
	userID string,
	roomID string,
) error {
	room, err := c.cache.GetChat(userID, roomID)
	if err != nil {
		return err
	}

	isAdmin, err := c.cache.IsAdmin(userID, roomID)
	if err != nil {
		return err
	}

	if room.Type == models.ChatType_GROUP && isAdmin != true {
		return status.Error(codes.PermissionDenied, "you are not group admin")
	}

	return c.cache.DeleteRoom(roomID)
}
