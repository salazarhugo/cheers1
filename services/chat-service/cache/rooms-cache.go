package cache

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"google.golang.org/api/chat/v1"
)

type RoomCache interface {
	SetMessage(msg *chat.Message)
	GetMessages(roomId string) []*pb.Message
	LikeMessage(roomId string, messageId string)
	UnlikeMessage(roomId string, messageId string)
	CreateGroup(name string, UUIDs []string) *pb.Room
	GetOrCreateDirectRoom(userId string, otherUserId string) *pb.Room
	LeaveRoom(userId string, roomId string)
	IsMember(userId string, roomId string) bool
	DeleteRoom(roomId string)
	GetRoomWithId(userId string, roomId string) *pb.Room
	GetRooms(userId string) []*pb.Room
	GetRoomMembers(roomId string) []string
	GetUser(userId string) (interface{}, error)
	DeleteTokens(userId string) int64
	AddToken(userId string, token string)
	GetTokens(userId string) []string
	GetOtherUserId(roomId string, userId string) (string, error)
	SetSeen(roomId string, userId string)
	GetRoomStatus(roomId string, userId string, otherUserId string) pb.RoomStatus
}
