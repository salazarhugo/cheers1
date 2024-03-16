package models

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (c *Chat) ToRoomPb(user *models.User) *pb.Room {
	return &pb.Room{
		Id:              c.Id,
		Name:            user.Name,
		Verified:        user.Verified,
		Typing:          false,
		Owner:           "",
		Type:            c.Type.ToRoomTypePb(),
		Status:          0,
		Admins:          c.Admins,
		Members:         nil,
		LastMessageText: c.LastMessageText,
		Picture:         user.Picture,
		LastMessageType: pb.MessageType_TEXT,
		CreateTime:      0,
		LastMessageTime: 0,
		LastMessageSeen: false,
		Archived:        false,
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
