package models

type ChatMessage struct {
	Id        string            `json:"id"`
	ChatId    string            `json:"chatId"`
	UserId    string            `json:"userId"`
	Text      string            `json:"text"`
	CreatedAt int64             `json:"createdAt"`
	IsSender  bool              `json:"isSender"`
	Status    ChatMessageStatus `json:"status"`
}
