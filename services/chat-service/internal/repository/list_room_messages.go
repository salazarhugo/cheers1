package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (c chatRepository) ListRoomMessages(
	roomID string,
	userID string,
) ([]*chat.MessageItem, error) {
	items := make([]*chat.MessageItem, 0)

	messages := c.cache.ListMessage(roomID, 10)
	for _, msg := range messages {
		msg.Status = chat.Message_DELIVERED
		items = append(items, &chat.MessageItem{
			Message: msg,
			Sender:  userID == msg.SenderId,
			Liked:   false,
		})
	}

	return items, nil
}
