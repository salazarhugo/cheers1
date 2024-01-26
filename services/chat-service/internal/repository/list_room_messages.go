package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (c chatRepository) ListRoomMessages(
	roomID string,
	viewerID string,
	page int,
	pageSize int,
) ([]*chat.MessageItem, error) {
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)
	items := make([]*chat.MessageItem, 0)

	messages := c.cache.ListMessage(roomID, offset, limit)
	for _, msg := range messages {
		msg.Status = chat.Message_DELIVERED
		items = append(items, &chat.MessageItem{
			Message: msg,
			Sender:  viewerID == msg.SenderId,
			Liked:   false,
		})
	}

	return items, nil
}
