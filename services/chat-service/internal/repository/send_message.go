package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"log"
	"strings"
	"time"
)

func (c chatRepository) SendMessage(
	messageID string,
	senderID string,
	roomID string,
	text string,
	replyToMessageId string,
) (*models.ChatMessage, error) {
	ctx := context.Background()

	msg := &models.ChatMessage{
		Id:        messageID,
		Text:      text,
		ChatId:    roomID,
		UserId:    senderID,
		CreatedAt: time.Now().UnixMilli(),
	}

	err := ValidateMessage(msg)
	if err != nil {
		return nil, err
	}

	// Store chat message into redis
	err = c.cache.CreateMessage(msg)
	if err != nil {
		return nil, err
	}

	websocketMsg := &models.WebSocketMessage{
		Type:        models.MESSAGE,
		UserId:      senderID,
		ChatMessage: *msg,
	}

	bytes, err := json.Marshal(websocketMsg)

	err = c.cache.SetLastRead(msg.ChatId, senderID, msg.CreatedAt)

	// Redis Pub/Sub
	err = c.cache.Publish(ctx, msg.ChatId, string(bytes)).Err()
	if err != nil {
		log.Println(err)
	}

	return msg, nil
}

func ValidateMessage(msg *models.ChatMessage) error {
	if msg == nil {
		return errors.New("nil message")
	}

	blank := strings.TrimSpace(msg.Text) == ""
	if blank {
		return errors.New("message text is blank")
	}

	return nil
}
