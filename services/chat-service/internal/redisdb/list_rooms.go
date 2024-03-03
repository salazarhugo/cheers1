package redisdb

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (cache *redisCache) ListRooms(userId string) ([]*pb.Room, error) {
	client := cache.client
	values, err := client.SMembers(
		context.Background(),
		getKeyUserRooms(userId),
	).Result()
	if err != nil {
		return nil, err
	}

	var rooms []*pb.Room
	for _, roomId := range values {
		room, err := cache.GetRoomWithId(userId, roomId)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room.ToRoomPb(nil))
	}

	return rooms, nil
}
