package notifications

func NewChatMessageNotification(
	username string,
	picture string,
	roomId string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "sent you a message",
		"avatar": picture,
		"roomId": roomId,
	}
}
