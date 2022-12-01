package models

type ChatMessage struct {
	RoomId string `json:"roomId"`
	Text   string `json:"text"`
}
