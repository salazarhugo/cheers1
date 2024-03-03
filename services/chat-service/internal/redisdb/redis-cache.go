package redisdb

import (
	"context"
	"fmt"
	"log"
)

const (
	keyRoomUsers           = "roomUsers"
	keyRoom                = "room"
	keyRoomSeen            = "roomSeen"
	keyRoomMessages        = "roomMessages"
	keyRoomMembers         = "roomMembers"
	keyUserRooms           = "userRooms"
	keyUserTokens          = "userTokens"
	keyRoomSenderRecipient = "roomSenderRecipient"
)

type User struct {
	Picture  string `json:"picture"`
	Username string `json:"username"`
}

func (cache *redisCache) SetSeen(roomId string, userId string) {
	lastMessage, err := getLatestMessage(cache.client, roomId)
	if err != nil {
		return
	}

	m := map[string]string{
		userId: lastMessage.Id,
	}

	cache.client.HSet(
		context.Background(),
		getKeyRoomSeen(roomId),
		m,
	)
}

func (cache *redisCache) DeleteTokens(userId string) int64 {
	client := cache.client
	successCount, err := client.Del(context.Background(), GetKeyUserTokens(userId)).Result()

	if err != nil {
		log.Println(err)
	}

	return successCount
}

func (cache *redisCache) GetTokens(userId string) []string {
	client := cache.client
	tokens, err := client.SMembers(context.Background(), GetKeyUserTokens(userId)).Result()
	if err != nil {
		return nil
	}

	return tokens
}

func (cache *redisCache) AddToken(userId string, token string) {
	client := cache.client
	client.SAdd(
		context.Background(),
		GetKeyUserTokens(userId),
		token,
	)
}

func (cache *redisCache) GetOtherUserId(
	roomId string,
	userId string,
) (string, error) {
	members, err := cache.client.SMembers(
		context.Background(),
		getKeyRoomMembers(roomId),
	).Result()
	if err != nil {
		return "", err
	}

	if members[0] == userId {
		return members[1], nil
	}
	return members[0], nil
}

func getKeyRoomSeen(roomUUID string) string {
	return fmt.Sprintf("%s:%s", keyRoomSeen, roomUUID)
}

func GetKeyUserTokens(userUUID string) string {
	return fmt.Sprintf("%s:%s", keyUserTokens, userUUID)
}

func getKeyUserRooms(userUUID string) string {
	return fmt.Sprintf("%s:%s", keyUserRooms, userUUID)
}

func getKeyRoomMembers(roomUUID string) string {
	return fmt.Sprintf("%s:%s", keyRoomMembers, roomUUID)
}

func getKeyRoomMessages(roomUUID string) string {
	return fmt.Sprintf("%s:%s", keyRoomMessages, roomUUID)
}

func getKeyRoom(roomUUID string) string {
	return fmt.Sprintf("%s:%s", keyRoom, roomUUID)
}

func getDirectRoomId(userId1 string, userId2 string) string {
	minUserId := userId1
	maxUserId := userId2

	if userId2 < userId1 {
		minUserId = userId2
		maxUserId = userId1
	}

	return fmt.Sprintf("%s%s", minUserId, maxUserId)
}
