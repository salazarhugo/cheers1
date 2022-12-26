package cache

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

const (
	keyUser                = "user"
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

func (cache *redisCache) IsMember(userId string, roomId string) bool {
	isMember, err := cache.client.SIsMember(context.Background(), getKeyRoomMembers(roomId), userId).Result()

	if err != nil {
		panic(err)
	}

	return isMember
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

	memberCount, _ := client.SCard(ctx, getKeyRoomMembers(roomId)).Result()

	if memberCount == 1 {
		cache.DeleteRoom(roomId)
	}
}

func (cache *redisCache) DeleteRoom(roomId string) error {
	var ctx = context.Background()
	client := cache.client

	members, err := client.SMembers(ctx, getKeyRoomMembers(roomId)).Result()
	if err != nil {
		return err
	}

	for _, uid := range members {
		// Delete user rooms
		client.SRem(ctx, getKeyUserRooms(uid), roomId)
	}

	// Delete room
	client.Del(ctx, getKeyRoom(roomId))
	// Delete room seen
	client.Del(ctx, getKeyRoomSeen(roomId))
	// Delete room members
	client.Del(ctx, getKeyRoomMembers(roomId))
	// Delete room messages
	client.Del(ctx, getKeyRoomMessages(roomId))

	return nil
}

func (cache *redisCache) SetSeen(roomId string, userId string) {
	lastMessage := getLatestMessage(cache.client, roomId)
	if lastMessage == nil {
		return
	}
	m := map[string]string{userId: lastMessage.Id}
	cache.client.HSet(context.Background(), getKeyRoomSeen(roomId), m)
}

func (cache *redisCache) GetRoomStatus(roomId string, userId string, otherUserId string) pb.RoomStatus {

	lastSeenMsg, err := cache.client.HMGet(context.Background(), getKeyRoomSeen(roomId), userId, otherUserId).Result()
	if err != nil {
		log.Println(err)
	}

	lastMsg := getLatestMessage(cache.client, roomId)
	if lastMsg == nil {
		return pb.RoomStatus_EMPTY
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
	return status
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
		panic(err)
	}

	return tokens
}

func (cache *redisCache) AddToken(userId string, token string) {
	client := cache.client
	client.SAdd(context.Background(), GetKeyUserTokens(userId), token)
}

func (cache *redisCache) ListMessage(roomId string, pageSize int64) []*pb.Message {
	client := cache.client

	values, err := client.ZRevRange(context.Background(), getKeyRoomMessages(roomId), 0, pageSize).Result()
	if err != nil {
		panic(err)
	}
	var messages []*pb.Message

	for i := range values {
		message := &pb.Message{}
		dec := json2.NewDecoder(strings.NewReader(values[i]))
		err := dec.Decode(message)
		if err != nil {
			panic(err)
		}
		message.Status = pb.Message_DELIVERED
		messages = append(messages, message)
	}

	return messages
}

func (cache *redisCache) SetMessage(msg *pb.Message) error {
	client := cache.client

	buff := bytes.NewBufferString("")
	enc := json2.NewEncoder(buff)
	err := enc.Encode(msg)
	if err != nil {
		return err
	}

	err = client.Watch(context.Background(), func(tx *redis.Tx) error {
		cmd := tx.ZAdd(context.Background(), getKeyRoomMessages(msg.GetRoomId()), redis.Z{
			Score:  float64(msg.CreateTime),
			Member: buff.String(),
		})

		if cmd.Err() != nil {
			return cmd.Err()
		}

		cmd2 := tx.Expire(context.Background(), getKeyRoomMessages(msg.GetRoomId()), 24*time.Hour)
		if cmd2.Err() != nil {
			return cmd.Err()
		}

		return nil
	})

	return err
}

func (cache *redisCache) ListRoomWithMessages(userId string) []*pb.RoomWithMessages {
	client := cache.client
	values, err := client.SMembers(context.Background(), getKeyUserRooms(userId)).Result()
	if err != nil {
		panic(err)
	}

	var rooms []*pb.RoomWithMessages
	for _, roomId := range values {
		room, _ := cache.GetRoomWithId(userId, roomId)
		messages := cache.ListMessage(roomId, 10)
		rooms = append(rooms, &pb.RoomWithMessages{
			Room:     room,
			Messages: messages,
		})
	}
	return rooms
}

func (cache *redisCache) GetOtherUserId(
	roomId string,
	userId string,
) (string, error) {
	members, err := cache.client.SMembers(context.Background(), getKeyRoomMembers(roomId)).Result()
	if err != nil {
		return "", err
	}

	if members[0] == userId {
		return members[1], nil
	}
	return members[0], nil
}

func (cache *redisCache) GetUserItem(
	userId string,
) (*user.UserItem, error) {
	res, err := cache.client.HGetAll(context.Background(), getKeyUser(userId)).Result()
	if err != nil {
		return nil, err
	}
	item := &user.UserItem{
		Id:          res["id"],
		Name:        res["name"],
		Username:    res["username"],
		Verified:    res["verified"] == "1",
		Picture:     res["picture"],
		HasFollowed: false,
		StoryState:  0,
	}

	return item, nil
}

func getKeyUser(userUUID string) string {
	return fmt.Sprintf("%s:%s", keyUser, userUUID)
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

func (cache *redisCache) CreateGroup(name string, UUIDs []string) *pb.Room {
	client := cache.client

	ctx := context.Background()

	room := &pb.Room{
		Id:              uuid.New().String(),
		LastMessageTime: 0,
		Name:            name,
		Status:          pb.RoomStatus_EMPTY,
		Type:            pb.RoomType_GROUP,
		Owner:           UUIDs[0],
	}

	cache.CreateRoom(room.Id, room)

	for _, uid := range UUIDs {
		client.SAdd(ctx, getKeyUserRooms(uid), room.Id)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), uid)
	}

	return room
}

func (cache *redisCache) GetOrCreateDirectRoom(
	userId string,
	otherUserId string,
) (*pb.Room, error) {
	client := cache.client

	ctx := context.Background()
	roomId := getDirectRoomId(userId, otherUserId)

	exists, err := client.Exists(context.Background(), getKeyRoom(roomId)).Result()
	if err != nil {
		return nil, err
	}

	if exists == 1 {
		room, err := cache.GetRoomWithId(userId, roomId)
		if err != nil {
			return nil, err
		}
		return room, nil
	}
	if exists == 0 {
		room := pb.Room{
			Id:              roomId,
			LastMessageTime: 0,
			Status:          pb.RoomStatus_EMPTY,
			Type:            pb.RoomType_DIRECT,
			Admins:          []string{userId, otherUserId},
		}
		cache.CreateRoom(roomId, &room)
		client.SAdd(ctx, getKeyUserRooms(userId), roomId)
		client.SAdd(ctx, getKeyUserRooms(otherUserId), roomId)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), userId, otherUserId)

		room_, err := cache.GetRoomWithId(userId, roomId)
		if err != nil {
			return nil, err
		}
		return room_, nil
	}

	return nil, nil
}

func getLatestMessage(client *redis.Client, roomId string) *pb.Message {
	lastMessage, err := client.ZRevRangeByScore(context.Background(), getKeyRoomMessages(roomId), &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}).Result()

	if len(lastMessage) == 0 {
		return nil
	}

	message := &pb.Message{}
	dec := json2.NewDecoder(strings.NewReader(lastMessage[0]))
	err = dec.Decode(message)
	if err != nil {
		panic(err)
	}

	return message
}

func (cache *redisCache) CreateRoom(roomId string, room *pb.Room) {
	client := cache.client

	json, err := json2.Marshal(room)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), getKeyRoom(roomId), json, cache.expires*time.Second)
}

func (cache *redisCache) GetRoomWithId(
	userId string,
	roomId string,
) (*pb.Room, error) {
	client := cache.client

	val, err := client.Get(context.Background(), getKeyRoom(roomId)).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Room %s not found", roomId))
	}

	if err != nil {
		return nil, err
	}

	room := pb.Room{}
	err = json2.Unmarshal([]byte(val), &room)
	if err != nil {
		panic(err)
	}

	otherUserId, err := cache.GetOtherUserId(roomId, userId)

	if room.Type == pb.RoomType_DIRECT {
		user, err := cache.GetUserItem(otherUserId)

		if err != nil {
			return nil, err
		}

		if user == nil {
			room.Name = "User not found"
		} else {
			room.Name = user.Name
			room.Picture = user.Picture
			room.Verified = user.Verified
		}
	}

	lastMessage := getLatestMessage(client, roomId)

	if lastMessage != nil {
		room.LastMessageText = lastMessage.Text
		room.LastMessageTime = lastMessage.CreateTime
		room.LastMessageType = lastMessage.Type
	}

	room.Status = cache.GetRoomStatus(roomId, userId, otherUserId)
	room.Members = cache.GetRoomMembers(roomId)

	return &room, nil
}
