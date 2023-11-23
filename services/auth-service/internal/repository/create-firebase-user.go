package repository

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (a *authRepository) CreateFirebaseUser() (*auth.UserRecord, error) {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	user, err := client.CreateUser(ctx, &auth.UserToCreate{})
	if err != nil {
		return nil, err
	}

	return user, nil
}
