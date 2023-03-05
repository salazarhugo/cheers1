package notifications

func AcceptedFriendRequestNotification(
	username string,
	picture string,
) map[string]string {
	return map[string]string{
		"title":  username,
		"body":   "accepted your friend request",
		"avatar": picture,
	}
}
