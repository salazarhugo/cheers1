package models

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (c *Chat) ToRoomPb(user *models.User) *pb.Room {
	return &pb.Room{
		Id:                 c.Id,
		Name:               user.Name,
		Verified:           user.Verified,
		Typing:             false,
		Owner:              "",
		Type:               c.Type.ToRoomTypePb(),
		Status:             ToChatStatus(c.Status),
		Admins:             c.Admins,
		Members:            nil,
		LastMessageText:    c.LastMessageText,
		Picture:            user.Picture,
		LastMessageType:    pb.MessageType_TEXT,
		CreateTime:         0,
		LastMessageTime:    c.LastMessageTime,
		LastMessageSeen:    false,
		Archived:           false,
		UnreadMessageCount: c.UnreadCount,
	}
}

func ToChatStatus(status string) pb.RoomStatus {
	switch status {
	case "EMPTY":
		return pb.RoomStatus_EMPTY
	case "NEW":
		return pb.RoomStatus_NEW
	case "SENT":
		return pb.RoomStatus_SENT
	case "OPENED":
		return pb.RoomStatus_OPENED
	case "RECEIVED":
		return pb.RoomStatus_RECEIVED
	default:
		return pb.RoomStatus_EMPTY
	}
}

func (c ChatType) ToRoomTypePb() pb.RoomType {
	switch c {
	case ChatType_DIRECT:
		return pb.RoomType_DIRECT
	case ChatType_GROUP:
		return pb.RoomType_GROUP
	default:
		return pb.RoomType_DIRECT
	}
}
