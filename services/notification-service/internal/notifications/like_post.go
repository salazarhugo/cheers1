package notifications

func LikePostNotification(
	username string,
	picture string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "liked your post",
		"avatar": picture,
	}
}
