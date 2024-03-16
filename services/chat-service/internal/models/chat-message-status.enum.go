package models

import (
	"encoding/json"
	"strings"
)

type ChatMessageStatus int32

const (
	ChatMessage_SENT      ChatMessageStatus = iota
	ChatMessage_DELIVERED ChatMessageStatus = iota
	ChatMessage_READ      ChatMessageStatus = iota
	ChatMessage_UNKNOWN   ChatMessageStatus = iota
)

// String - Creating common behavior - give the type a String function
func (w ChatMessageStatus) String() string {
	switch w {
	case ChatMessage_SENT:
		return "SENT"
	case ChatMessage_DELIVERED:
		return "DELIVERED"
	case ChatMessage_READ:
		return "READ"
	}

	return "UNKNOWN"
}

func (w ChatMessageStatus) EnumIndex() int {
	return int(w)
}

func (w ChatMessageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}

func (w *ChatMessageStatus) UnmarshalJSON(data []byte) (err error) {
	var suits string
	if err := json.Unmarshal(data, &suits); err != nil {
		return err
	}

	if *w, err = ParseChatMessageStatus(suits); err != nil {
		return err
	}

	return nil
}

func ParseChatMessageStatus(s string) (ChatMessageStatus, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "sent":
		return ChatMessage_SENT, nil
	case "delivered":
		return ChatMessage_DELIVERED, nil
	case "read":
		return ChatMessage_READ, nil

	}

	return ChatMessage_UNKNOWN, nil
}
