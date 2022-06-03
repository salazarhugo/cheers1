package main

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"log"
	"sync"
)

// Notification Single user notification to all of his devices
type Notification struct {
	ReceiverId string   `structs:"receiverId,omitempty"`
	Tokens     []string `structs:"tokens,omitempty"`
	Title      string   `structs:"title,omitempty"`
	Body       string   `structs:"body,omitempty"`
	Avatar     string   `structs:"avatar,omitempty"`
}

func SendNotification(notification *Notification, wg *sync.WaitGroup) {
	defer wg.Done()
	app := initializeAppDefault()
	a, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	br, err := a.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Data: map[string]string{
			"type":   "newMessage",
			"title":  notification.Title,
			"body":   notification.Body,
			"avatar": notification.Avatar,
		},
		Tokens: notification.Tokens,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Successfully sent notification")

	if br.FailureCount == 0 {
		return
	}

	var stillRegisteredTokens []string
	copy(stillRegisteredTokens, notification.Tokens)

	for idx, resp := range br.Responses {
		if resp.Success {
			continue
		}
		for i, token := range stillRegisteredTokens {
			if token == notification.Tokens[idx] {
				stillRegisteredTokens = remove(stillRegisteredTokens, i)
			}
		}
	}

	err = setRegistrationToken(notification.ReceiverId, stillRegisteredTokens)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func setRegistrationToken(userId string, registrationTokens []string) error {
	session := getSession(getDriver())
	defer session.Close()

	cypher := `MATCH (u:User {id: $userId})
				SET u.registrationTokens = $registrationTokens`

	params := map[string]interface{}{
		"userId":             userId,
		"registrationTokens": registrationTokens,
	}

	_, err := session.Run(cypher, params)

	return err
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
