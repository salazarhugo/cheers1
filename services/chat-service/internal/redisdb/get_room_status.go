package redisdb

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (cache *redisCache) GetRoomStatus(
	chatID string,
	userId string,
	otherUserId string,
) (pb.RoomStatus, error) {
	lastReadOfViewer, err := cache.GetLastRead(chatID, userId)
	lastReadOfOtherUser, err := cache.GetLastRead(chatID, otherUserId)

	lastMsg, err := getLatestMessage(cache.client, chatID)
	if lastMsg == nil {
		return pb.RoomStatus_EMPTY, err
	}

	isLastMessageMe := lastMsg.UserId == userId
	seenByMe := lastReadOfViewer > lastMsg.CreatedAt
	seenByOther := lastReadOfOtherUser > lastMsg.CreatedAt

	status := pb.RoomStatus_EMPTY

	if isLastMessageMe {
		if seenByOther {
			status = pb.RoomStatus_OPENED
		} else {
			status = pb.RoomStatus_SENT
		}
	} else {
		if seenByMe {
			status = pb.RoomStatus_RECEIVED
		} else {
			status = pb.RoomStatus_NEW
		}
	}

	return status, nil
}
