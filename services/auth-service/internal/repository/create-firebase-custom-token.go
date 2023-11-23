package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (a *authRepository) CreateFirebaseCustomToken(userID string) (string, error) {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return "", err
	}

	token, err := client.CustomToken(ctx, userID)
	if err != nil {
		return "", err
	}

	return token, nil
}
