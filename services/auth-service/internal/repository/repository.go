package repository

import (
	"firebase.google.com/go/v4/auth"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
	"gorm.io/gorm"
	"os"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*User, error)
	CheckUsername(username string) (bool, error)
	GetUserCredentials(username string) ([]*Credential, error)
	GetAuthnUser(username string) (*AuthnUser, error)
	CreateUser(user *User) error
	CreateFirebaseUser() (*auth.UserRecord, error)
	CreateFirebaseCustomToken(userID string) (string, error)
	GetOrCreateUser(username string) (*AuthnUser, error)
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

func (a *authRepository) CheckUsername(username string) (bool, error) {
	db := a.spanner
	var exists bool
	result := db.Raw("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if result.Error != nil {
		return exists, result.Error
	}

	return exists, nil
}

func (a *authRepository) GetAuthnUser(username string) (*AuthnUser, error) {
	// Get the user
	user, err := a.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Get user credentials
	credentials, err := a.GetUserCredentials(username)
	if err != nil {
		return nil, err
	}

	authnUser := user.ToAuthnUser()
	for _, credential := range credentials {
		authnUser.AddCredential(*credential.ToWebAuthnCredential())
	}

	return authnUser, nil
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
