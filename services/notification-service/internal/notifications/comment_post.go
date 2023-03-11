package notifications

func CommentPostNotification(
	username string,
	picture string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "commented on your post.",
		"avatar": picture,
	}
}
