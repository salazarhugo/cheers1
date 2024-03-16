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

	messages, err := c.cache.ListChatMessage(
		roomID,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	otherUserID, err := c.cache.GetOtherUserId(roomID, viewerID)
	lastReadAt, err := c.cache.GetLastRead(roomID, otherUserID)

	for _, msg := range messages {
		msgPb := msg.ToChatMessagePb()

		if lastReadAt >= msg.CreatedAt {
			msgPb.Status = chat.Message_READ
		} else {
			msgPb.Status = chat.Message_DELIVERED
		}

		items = append(items, &chat.MessageItem{
			Message: msgPb,
			Sender:  viewerID == msg.UserId,
			Liked:   false,
		})
	}

	return items, nil
}
