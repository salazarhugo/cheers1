package models

type Chat struct {
	Id              string   `json:"id"`
	Message         string   `json:"message"`
	Name            string   `json:"name"`
	LastMessageTime int64    `json:"last_message_time"`
	LastMessageText string   `json:"last_message_text"`
	LastMessageType string   `json:"last_message_type"`
	Type            ChatType `json:"type"`
	Status          string   `json:"status"`
	Admins          []string `json:"admins"`
}

type ChatType int32

const (
	ChatType_DIRECT ChatType = iota
	ChatType_GROUP  ChatType = iota
)

// String - Creating common behavior - give the type a String function
func (w ChatType) String() string {
	return [...]string{"Direct", "Group"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w ChatType) EnumIndex() int {
	return int(w)
}
