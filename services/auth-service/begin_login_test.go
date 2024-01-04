package auth_service

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/app"
	"log"
	"testing"
)

func TestBeginLogin(t *testing.T) {
	server := app.NewServer()
	response, err := server.BeginLogin(context.Background(), &auth.BeginLoginRequest{Username: "salazar"})
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(response)
}
