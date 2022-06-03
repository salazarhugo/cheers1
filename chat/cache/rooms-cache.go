package cache

import "chat/chatpb"

type RoomCache interface {
	SetMessage(msg *chatpb.Message)
	GetMessages(roomId string) []*chatpb.Message
	CreateGroup(name string, UUIDs []string) *chatpb.Room
	GetOrCreateDirectRoom(userId string, otherUserId string) *chatpb.Room
	LeaveRoom(userId string, roomId string)
	DeleteRoom(roomId string)
	GetRoomWithId(userId string, roomId string) *chatpb.Room
	GetRooms(userId string) []*chatpb.Room
	GetRoomMembers(roomId string) []string
	GetUser(userId string) (interface{}, error)
	AddToken(userId string, token string)
	GetTokens(userId string) []string
	GetOtherUserId(roomId string, userId string) (string, error)
	SetSeen(roomId string, userId string)
	GetRoomStatus(roomId string, userId string, otherUserId string) chatpb.RoomStatus
}
