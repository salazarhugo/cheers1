package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func (c chatRepository) CreateRoom(
	name string,
	members []string,
) (*pb.Room, error) {
	// Direct Chat
	if len(members) == 2 {
		chat, err := c.cache.GetOrCreateDirectRoom(members[0], members[1])
		if err != nil {
			return nil, err
		}

		if chat.Type == models.ChatType_DIRECT {
			userID, err := c.cache.GetOtherUserId(chat.Id, members[0])
			user, err := c.GetUserNode(userID)
			if err != nil {
				return nil, err
			}
			room := chat.ToRoomPb(user)
			return room, nil
		}
	}

	room, err := c.cache.CreateGroup(name, members)
	if err != nil {
		return nil, err
	}

	return room.ToRoomPb(nil), nil
}
