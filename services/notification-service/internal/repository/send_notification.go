package repository

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (r repository) SendNotification(
	userWithToken map[string][]string,
	data map[string]string,
) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	fcmClient, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	for userId, tokens := range userWithToken {
		response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
			Data:   data,
			Tokens: tokens,
		})
		if err != nil {
			log.Println(err)
			continue
		}

		go r.RemoveExpiredTokens(userId, tokens, response.Responses)

		log.Println("Response success count : ", response.SuccessCount)
		log.Println("Response failure count : ", response.FailureCount)
	}
	return nil
}
