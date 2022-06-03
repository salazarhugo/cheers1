package main

import (
	"context"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/users/available/:username" {
			return next(c)
		}
		app := initializeAppDefault()
		client := initializeAuthClient(app)
		authorizationToken := c.Request().Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			log.Println("Empty Authorization Token")
			return c.JSON(http.StatusBadRequest, "Empty token")
		}

		// Verify Id token
		token, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
		if err != nil {
			log.Println("Invalid Authorization Token or revoked or disabled")
			return c.String(http.StatusBadRequest, fmt.Sprint("Invalid token: ", idToken))
		}
		c.Set("userId", token.UID)
		return next(c)
	}
}

// GET("/token", createCustomToken)
func createCustomToken(c echo.Context) error {
	cc := c.(*CustomContext)

	app := initializeAppDefault()
	client := initializeAuthClient(app)

	username := cc.QueryParam("username")

	claims := map[string]interface{}{
		//"premiumAccount": true,
	}

	token, err := client.CustomTokenWithClaims(context.Background(), username, claims)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	return cc.String(http.StatusOK, token)
}
