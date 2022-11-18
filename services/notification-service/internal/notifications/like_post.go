package notifications

import (
	"firebase.google.com/go/v4/messaging"
	"fmt"
)

func LikePostNotification(
	username string,
	picture string,
) messaging.Notification {
	return messaging.Notification{
		Title:    "Cheers",
		Body:     fmt.Sprintf("%s liked your post", username),
		ImageURL: picture,
	}
}
