package auth

import (
	"context"
	"firebase.google.com/go/v4"
	"github.com/labstack/gommon/log"
)

func InitializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error utilsializing app: %v\n", err)
	}

	return app
}
