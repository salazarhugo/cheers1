package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (c chatRepository) ListRoomMessages(
	roomID string,
	request *chat.ListRoomMessagesRequest,
) ([]*chat.Message, error) {
	messages := c.cache.GetMessages(roomID)
	return messages, nil
}
