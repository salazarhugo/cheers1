package repository

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (r repository) SendNotification(
	userWithToken map[string][]string,
	notification messaging.Notification,
) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	fcmClient, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	for userId, tokens := range userWithToken {
		response, _ := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
			Notification: &notification,
			Tokens:       tokens,
		})
		go r.RemoveExpiredTokens(userId, tokens, response.Responses)

		log.Println("Response success count : ", response.SuccessCount)
		log.Println("Response failure count : ", response.FailureCount)
	}
	return nil
}
