package models

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (c *ChatMessage) ToChatMessagePb() *pb.Message {
	return &pb.Message{
		Id:             c.Id,
		RoomId:         c.ChatId,
		Text:           c.Text,
		Picture:        "",
		SenderId:       c.UserId,
		SenderPicture:  "",
		SenderName:     "",
		SenderUsername: "",
		LikeCount:      0,
		Type:           pb.MessageType_TEXT,
		Status:         pb.Message_DELIVERED,
		CreateTime:     c.CreatedAt,
	}
}
