package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
	"gorm.io/gorm"
	"os"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
	GetOrCreateUser(username string) (*User, error)
	CreateAdmin(userID string) error
	CreateModerator(userID string) error
	CreateBusinessAccount(userID string) error
	VerifyUser(userID string) error
	CreateClaim(userID string, claim constants.Claim) error
	CreateCredential(credential *Credential) error
}

type authRepository struct {
	spanner *gorm.DB
	redis   *redis.Client
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
