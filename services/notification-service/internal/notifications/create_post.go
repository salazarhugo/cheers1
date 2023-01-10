package notifications

import "strings"

func CreatePostNotification(
	name string,
	username string,
	picture string,
	location string,
	drink string,
	caption string,
	isPhoto bool,
) map[string]string {
	sir := name
	if name == "" {
		sir = username
	}

	title := sir + " just posted a message."
	if isPhoto {
		title = sir + " just posted a photo."
	}
	if drink != "NONE" && drink != "" {
		title = sir + " is drinking " + strings.ToLower(drink)
	}

	body := caption

	if location != "" {
		body = "📍 " + location
	}

	return map[string]string{
		"title":  title,
		"body":   body,
		"avatar": picture,
	}
}
