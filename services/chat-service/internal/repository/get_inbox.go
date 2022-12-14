package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (c chatRepository) GetInbox(userID string) ([]*pb.RoomWithMessages, error) {
	rooms := c.cache.ListRoomWithMessages(userID)
	return rooms, nil
}
