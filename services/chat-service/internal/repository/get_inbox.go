package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (c chatRepository) GetInbox(viewerID string) ([]*pb.RoomWithMessages, error) {
	rooms, err := c.cache.ListRooms(viewerID)
	if err != nil {
		return nil, err
	}

	var inbox []*pb.RoomWithMessages
	for _, room := range rooms {
		messages, err := c.ListRoomMessages(room.Id, viewerID, 1, 10)
		if err != nil {
			continue
		}

		inbox = append(inbox, &pb.RoomWithMessages{
			Room:     room,
			Messages: messages,
		})
	}

	return inbox, nil
}
