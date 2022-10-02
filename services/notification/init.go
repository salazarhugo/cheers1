package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"log"
)

func initializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}
