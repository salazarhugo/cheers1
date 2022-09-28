package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"time"
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

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	start := time.Now()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}

	userInfoHeader, ok := md["x-apigateway-api-userinfo"]
	if !ok {
		log.Println("Empty Userinfo Header")
		return nil, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	encodedUser := userInfoHeader[0]

	decodedBytes, err := base64.RawURLEncoding.DecodeString(encodedUser)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid User Info")
	}

	decoder := json.NewDecoder(bytes.NewReader(decodedBytes))
	var userInfo UserInfo
	err = decoder.Decode(&userInfo)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid User Info")
	}

	log.Println(userInfo.UserID)
	ctx = metadata.AppendToOutgoingContext(ctx, "userId", userInfo.UserID)

	h, err := handler(ctx, req)

	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
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
