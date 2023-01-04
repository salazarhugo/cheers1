package notifications

func NewGroupMessageNotification(
	name string,
	picture string,
	roomId string,
) map[string]string {
	return map[string]string{
		"title":  name,
		"body":   "new chat",
		"avatar": picture,
		"roomId": roomId,
	}
}
