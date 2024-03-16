package models

type WebSocketMessage struct {
	Type        WsType      `json:"type"`
	UserId      string      `json:"userId"`
	IsViewer    bool        `json:"isViewer"`
	ChatMessage ChatMessage `json:"chatMessage"`
	Typing      Typing      `json:"typing"`
	Presence    Presence    `json:"presence"`
}

type Typing struct {
	ChatId   string `json:"chatId"`
	IsTyping bool   `json:"isTyping"`
}

type Presence struct {
	ChatId    string `json:"chatId"`
	IsPresent bool   `json:"isPresent"`
}
