package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	_ "firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserInfo struct {
	Name          string   `json:"name"`
	Picture       string   `json:"picture"`
	Iss           string   `json:"iss"`
	Aud           string   `json:"aud"`
	AuthTime      int      `json:"auth_time"`
	UserID        string   `json:"user_id"`
	Sub           string   `json:"sub"`
	Iat           int      `json:"iat"`
	Exp           int      `json:"exp"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"email_verified"`
	Firebase      Firebase `json:"firebase"`
}

type Identities struct {
	GoogleCom []string `json:"google.com"`
	Email     []string `json:"email"`
}

type Firebase struct {
	Identities     Identities `json:"identities"`
	SignInProvider string     `json:"sign_in_provider"`
}

func UserInfoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		encodedUser := c.Request().Header.Get("X-Apigateway-Api-Userinfo")
		if encodedUser == "" {
			log.Println("Empty Userinfo Header")
			return c.JSON(http.StatusForbidden, "Empty X-Apigateway-Api-Userinfo header")
		}
		decodedBytes, err := base64.RawURLEncoding.DecodeString(encodedUser)
		if err != nil {
			return c.JSON(http.StatusForbidden, "Invalid UserInfo")
		}
		decoder := json.NewDecoder(bytes.NewReader(decodedBytes))
		var userInfo UserInfo
		err = decoder.Decode(&userInfo)
		if err != nil {
			return c.JSON(http.StatusForbidden, "Invalid UserInfo")
		}

		c.Set("userId", userInfo.UserID)
		return next(c)
	}
}
