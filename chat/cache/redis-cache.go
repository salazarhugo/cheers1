package cache

import (
	"bytes"
	"chat/chatpb"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strings"
	"time"
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
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Username          string `json:"username"`
}

type redisCache struct {
	expires time.Duration
	client  *redis.Client
}

func (cache *redisCache) GetRoomMembers(roomId string) []string {
	members, err := cache.client.SMembers(context.Background(), getKeyRoomMembers(roomId)).Result()
	if err != nil {
		panic(err)
	}

	return members
}

func (cache *redisCache) LeaveRoom(userId string, roomId string) {
	var ctx = context.Background()
	client := cache.client

	client.SRem(ctx, getKeyUserRooms(userId), roomId)
	client.SRem(ctx, getKeyRoomMembers(roomId), userId)
}

func (cache *redisCache) DeleteRoom(roomId string) {
	var ctx = context.Background()
	client := cache.client

	members, err := client.SMembers(ctx, getKeyRoomMembers(roomId)).Result()
	if err != nil {
		panic(err)
	}

	for _, uid := range members {
		client.SRem(ctx, getKeyUserRooms(uid), roomId)
	}

	client.Del(ctx, getKeyRoom(roomId))
	client.Del(ctx, getKeyRoomSeen(roomId))
	client.Del(ctx, getKeyRoomMembers(roomId))
	client.Del(ctx, getKeyRoomMessages(roomId))
}

func (cache *redisCache) SetSeen(roomId string, userId string) {
	lastMessage := getLatestMessage(cache.client, roomId)
	if lastMessage == nil {
		return
	}
	m := map[string]string{userId: lastMessage.Id}
	cache.client.HSet(context.Background(), getKeyRoomSeen(roomId), m)
}

func (cache *redisCache) GetRoomStatus(roomId string, userId string, otherUserId string) chatpb.RoomStatus {

	lastSeenMsg, err := cache.client.HMGet(context.Background(), getKeyRoomSeen(roomId), userId, otherUserId).Result()
	if err != nil {
		log.Println(err)
	}

	lastMsg := getLatestMessage(cache.client, roomId)
	if lastMsg == nil {
		return chatpb.RoomStatus_EMPTY
	}
	isLastMessageMe := lastMsg.Sender == userId
	seenByMe := lastMsg.Id == lastSeenMsg[0]
	seenByOther := lastMsg.Id == lastSeenMsg[1]

	status := chatpb.RoomStatus_EMPTY

	if isLastMessageMe {
		if seenByOther {
			status = chatpb.RoomStatus_OPENED
		} else {
			status = chatpb.RoomStatus_SENT
		}
	} else {
		if seenByMe {
			status = chatpb.RoomStatus_RECEIVED
		} else {
			status = chatpb.RoomStatus_NEW
		}
	}
	return status
}

func (cache *redisCache) GetTokens(userId string) []string {
	client := cache.client
	tokens, err := client.SMembers(context.Background(), getKeyUserTokens(userId)).Result()
	if err != nil {
		panic(err)
	}

	return tokens
}

func (cache *redisCache) AddToken(userId string, token string) {
	client := cache.client
	client.SAdd(context.Background(), getKeyUserTokens(userId), token)
}

func (cache *redisCache) GetMessages(roomId string) []*chatpb.Message {
	client := cache.client

	values, err := client.ZRevRange(context.Background(), getKeyRoomMessages(roomId), 0, 50).Result()
	if err != nil {
		panic(err)
	}
	var messages []*chatpb.Message

	for i := range values {
		message := &chatpb.Message{}
		dec := json2.NewDecoder(strings.NewReader(values[i]))
		err := dec.Decode(message)
		if err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}

	return messages
}

func (cache *redisCache) SetMessage(msg *chatpb.Message) {
	client := cache.client

	buff := bytes.NewBufferString("")
	enc := json2.NewEncoder(buff)
	err := enc.Encode(msg)
	if err != nil {
		panic(err)
	}

	client.ZAdd(context.Background(), getKeyRoomMessages(msg.GetRoom().GetId()), &redis.Z{
		Score:  float64(msg.GetCreated().Seconds),
		Member: buff.String(),
	})
	client.Expire(context.Background(), getKeyRoomMessages(msg.GetRoom().GetId()), 24*time.Hour)
}

func (cache *redisCache) GetRooms(userId string) []*chatpb.Room {
	client := cache.client
	values, err := client.SMembers(context.Background(), getKeyUserRooms(userId)).Result()

	if err != nil {
		panic(err)
	}

	var rooms []*chatpb.Room
	for _, roomId := range values {
		room := cache.GetRoomWithId(userId, roomId)
		rooms = append(rooms, room)
	}
	return rooms
}

func (cache *redisCache) GetOtherUserId(roomId string, userId string) (string, error) {
	arr := strings.Split(roomId, ":")

	if len(arr) <= 1 {
		return "", bytes.ErrTooLarge
	}

	if arr[0] == userId {
		return arr[1], nil
	}
	return arr[0], nil
}

func (cache *redisCache) GetUser(userId string) (interface{}, error) {
	driver, err := neo4j.NewDriver(
		"neo4j+s://528a253a.databases.neo4j.io:7687",
		neo4j.BasicAuth("neo4j", "XRoQ6Lmz9QlFFTcwCWIWwR1o88hLfzV_HnP9mzDJuwc", ""))

	if err != nil {
		return nil, err
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	user, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId })
					RETURN properties(u)`,
			map[string]interface{}{"userId": userId})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getKeyRoomSeen(roomUUID string) string {
	return fmt.Sprintf("%s:%s", keyRoomSeen, roomUUID)
}

func getKeyUserTokens(userUUID string) string {
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

	return fmt.Sprintf("%s:%s", minUserId, maxUserId)
}

func (cache *redisCache) CreateGroup(name string, UUIDs []string) *chatpb.Room {
	client := cache.client

	ctx := context.Background()

	room := &chatpb.Room{
		Id:              uuid.New().String(),
		LastMessageTime: timestamppb.Now(),
		Name:            name,
		Status:          chatpb.RoomStatus_EMPTY,
		Type:            chatpb.RoomType_GROUP,
		Owner:           UUIDs[0],
	}

	cache.CreateRoom(room.Id, room)

	for _, uid := range UUIDs {
		client.SAdd(ctx, getKeyUserRooms(uid), room.Id)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), uid)
	}

	return room
}

func (cache *redisCache) GetOrCreateDirectRoom(userId string, otherUserId string) *chatpb.Room {
	client := cache.client

	ctx := context.Background()
	roomId := getDirectRoomId(userId, otherUserId)

	exists, err := client.Exists(context.Background(), getKeyRoom(roomId)).Result()
	if err != nil {
		panic(err)
	}

	if exists == 1 {
		room := cache.GetRoomWithId(userId, roomId)
		return room
	}
	if exists == 0 {
		room := chatpb.Room{
			Id:              roomId,
			LastMessageTime: timestamppb.Now(),
			Status:          chatpb.RoomStatus_EMPTY,
			Type:            chatpb.RoomType_DIRECT,
		}
		cache.CreateRoom(roomId, &room)
		client.SAdd(ctx, getKeyUserRooms(userId), roomId)
		client.SAdd(ctx, getKeyUserRooms(otherUserId), roomId)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), userId, otherUserId)

		room_ := cache.GetRoomWithId(userId, roomId)
		return room_
	}

	return nil
}

func getLatestMessage(client *redis.Client, roomId string) *chatpb.Message {
	lastMessage, err := client.ZRevRangeByScore(context.Background(), getKeyRoomMessages(roomId), &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}).Result()

	if len(lastMessage) == 0 {
		return nil
	}

	message := &chatpb.Message{}
	dec := json2.NewDecoder(strings.NewReader(lastMessage[0]))
	err = dec.Decode(message)
	if err != nil {
		panic(err)
	}

	return message
}

func (cache *redisCache) CreateRoom(roomId string, room *chatpb.Room) {
	client := cache.client

	json, err := json2.Marshal(room)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), getKeyRoom(roomId), json, cache.expires*time.Second)
}

func (cache *redisCache) GetRoomWithId(userId string, roomId string) *chatpb.Room {
	client := cache.client

	val, err := client.Get(context.Background(), getKeyRoom(roomId)).Result()
	if err != nil {
		return nil
	}

	room := chatpb.Room{}
	err = json2.Unmarshal([]byte(val), &room)
	if err != nil {
		panic(err)
	}

	otherUserId, err := cache.GetOtherUserId(roomId, userId)
	log.Println(otherUserId)

	if room.Type == chatpb.RoomType_DIRECT {
		user, err := cache.GetUser(otherUserId)

		if err != nil {
			panic(err)
		}

		if user == nil {
			room.Name = "User not found"
		} else {
			userMap := user.(map[string]interface{})

			name, ok := userMap["name"].(string)
			if ok {
				room.Name = name
			}
			username, ok := userMap["username"].(string)
			if ok {
				room.Username = username
			}
			pp, ok := userMap["profilePictureUrl"].(string)
			if ok {
				room.ProfilePictureUrl = pp
			}
			verified, ok := userMap["verified"].(bool)
			if ok {
				room.Verified = verified
			}
		}
	}

	lastMessage := getLatestMessage(client, roomId)

	if lastMessage != nil {
		room.LastMessageText = lastMessage.Message
		room.LastMessageTime = lastMessage.Created
		room.LastMessageType = lastMessage.Type
	}

	room.Status = cache.GetRoomStatus(roomId, userId, otherUserId)
	room.Members = cache.GetRoomMembers(roomId)

	return &room
}

func NewRedisCache(addr string, password string, db int, exp time.Duration) RoomCache {
	return &redisCache{
		expires: exp,
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}
