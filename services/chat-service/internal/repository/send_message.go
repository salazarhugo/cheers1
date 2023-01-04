package repository

import (
	"context"
	"errors"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	"strings"
	"time"
)

func (c chatRepository) SendMessage(
	messageID string,
	senderID string,
	roomID string,
	text string,
) (*chat.Message, error) {
	ctx := context.Background()

	msg := &chat.Message{
		Id:         messageID,
		Text:       text,
		RoomId:     roomID,
		SenderId:   senderID,
		CreateTime: time.Now().Unix(),
	}

	err := ValidateMessage(msg)
	if err != nil {
		return nil, err
	}

	// Store message in cache
	err = c.cache.SetMessage(msg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bytes, err := protojson.Marshal(msg)
	// Redis Pub/Sub
	c.cache.Publish(ctx, msg.RoomId, bytes)

	// Google Cloud Pub/Sub
	user, err := c.cache.GetUserItem(msg.SenderId)
	if err != nil {
		log.Println(err)
	}
	membersIds := c.cache.GetRoomMembers(msg.RoomId)
	if err != nil {
		log.Println(err)
	}
	members, err := c.cache.ListUser(membersIds)
	if err != nil {
		log.Println(err)
	}
	room, err := c.cache.GetRoomWithId(senderID, roomID)
	if err != nil {
		log.Println(err)
	}

	err = pubsub.PublishProtoWithBinaryEncoding(os.Getenv("EVENT_BUS_TOPIC"), &chat.ChatEvent{
		Event: &chat.ChatEvent_Create{
			Create: &chat.CreateMessage{
				Message: msg,
				Sender:  user,
				Members: members,
				Room:    room,
			},
		},
	})

	if err != nil {
		log.Println(err)
	}

	return msg, err
}

func ValidateMessage(msg *chat.Message) error {
	if msg == nil {
		return errors.New("nil message")
	}

	blank := strings.TrimSpace(msg.Text) == ""
	if blank {
		return errors.New("message text is blank")
	}

	return nil
}
