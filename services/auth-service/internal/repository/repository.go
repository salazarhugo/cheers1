package repository

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/model"
	"gorm.io/gorm"
	"os"
	"time"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	GetSession(username string) (*webauthn.SessionData, error)
	PutSession(username string, session *webauthn.SessionData) error
	CreateAdmin(userID string) error
	CreateModerator(userID string) error
	CreateBusinessAccount(userID string) error
	VerifyUser(userID string) error
	CreateClaim(userID string, claim constants.Claim) error
	CreateCredential(userID string, publicKey string) (string, error)
}

type authRepository struct {
	spanner *gorm.DB
	redis   *redis.Client
}

func (a authRepository) PutSession(
	username string,
	session *webauthn.SessionData,
) error {
	client := a.redis

	json, err := json2.Marshal(session)
	if err != nil {
		return err
	}

	err = client.Set(context.Background(), getKeySession(username), json, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func getKeySession(username string) string {
	return fmt.Sprintf("%s:%s", "session", username)
}

func (a authRepository) GetSession(
	username string,
) (*webauthn.SessionData, error) {
	client := a.redis

	session := webauthn.SessionData{}

	val, err := client.Get(context.Background(), getKeySession(username)).Result()
	err = json2.Unmarshal([]byte(val), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func NewRepository() AuthRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})
	return &authRepository{
		spanner: utils.GetSpanner(),
		redis:   rdb,
	}
}
