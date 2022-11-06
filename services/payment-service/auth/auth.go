package auth

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/salazarhugo/cheers1/services/payment-service/utils"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/handleStripeEvents" {
			return next(c)
		}
		if c.Path() == "/ticket/validate" {
			return next(c)
		}
		app := utils.InitializeAppDefault()
		client := utils.InitializeAuthClient(app)
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
			return c.String(http.StatusForbidden, fmt.Sprint("Invalid token: ", idToken))
		}
		c.Set("userId", token.UID)
		return next(c)
	}
}
