package cache

import "chat/chatpb"

type RoomCache interface {
	SetMessage(msg *chatpb.Message)
	GetMessages(roomId string) []*chatpb.Message

	LikeMessage(roomId string, messageId string)
	UnlikeMessage(roomId string, messageId string)

	CreateGroup(name string, UUIDs []string) *chatpb.Room
	GetOrCreateDirectRoom(userId string, otherUserId string) *chatpb.Room
	LeaveRoom(userId string, roomId string)
	IsMember(userId string, roomId string) bool
	DeleteRoom(roomId string)
	GetRoomWithId(userId string, roomId string) *chatpb.Room
	GetRooms(userId string) []*chatpb.Room
	GetRoomMembers(roomId string) []string
	GetUser(userId string) (interface{}, error)
	DeleteTokens(userId string) int64
	AddToken(userId string, token string)
	GetTokens(userId string) []string
	GetOtherUserId(roomId string, userId string) (string, error)
	SetSeen(roomId string, userId string)
	GetRoomStatus(roomId string, userId string, otherUserId string) chatpb.RoomStatus
}
