package notifications

func FollowUserNotification(
	username string,
	picture string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "started following you",
		"avatar": picture,
	}
}
