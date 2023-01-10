package notifications

func FriendRequestNotification(
	username string,
	picture string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "wants to be your friend",
		"avatar": picture,
	}
}
