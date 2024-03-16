package main

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
	"time"
)

func ValidateJwt(token string) (*auth.Token, error) {
	// Initialize Firebase app
	ctx := context.Background()
	app := utils.InitializeAppDefault()

	// Get auth client
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
		return nil, err
	}

	// Verify ID token
	idToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		log.Printf("Error verifying ID token: %v\n", err)
		return nil, err
	}

	// Check if token is valid
	if idToken.Expires < time.Now().Unix() {
		log.Printf("Token has expired\n")
		return nil, err
	}

	// Token is valid
	return idToken, nil
}
