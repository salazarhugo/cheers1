package repository

import (
	"firebase.google.com/go/v4/messaging"
	"log"
)

func (r repository) RemoveExpiredTokens(
	userId string,
	tokens []string,
	responses []*messaging.SendResponse,
) {
	for i, item := range responses {
		if item.Success {
			continue
		}
		token := tokens[i]
		err := r.DeleteToken(userId, token)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted token: %s", token)
	}
}
