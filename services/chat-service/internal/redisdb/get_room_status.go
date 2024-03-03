package redisdb

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (cache *redisCache) GetRoomStatus(
	roomId string,
	userId string,
	otherUserId string,
) (pb.RoomStatus, error) {
	lastSeenMsg, err := cache.client.HMGet(
		context.Background(),
		getKeyRoomSeen(roomId),
		userId,
		otherUserId,
	).Result()
	if err != nil {
		return 0, err
	}

	lastMsg, err := getLatestMessage(cache.client, roomId)
	if lastMsg == nil {
		return pb.RoomStatus_EMPTY, err
	}

	isLastMessageMe := lastMsg.SenderId == userId
	seenByMe := lastMsg.Id == lastSeenMsg[0]
	seenByOther := lastMsg.Id == lastSeenMsg[1]

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
