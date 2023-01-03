package notifications

func CreatePostNotification(
	username string,
	picture string,
	drink string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "is drinking " + drink,
		"avatar": picture,
	}
}
