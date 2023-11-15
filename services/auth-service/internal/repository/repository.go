package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateAdmin(userID string) error
	CreateModerator(userID string) error
	CreateBusinessAccount(userID string) error
	VerifyUser(userID string) error
	CreateClaim(userID string, claim constants.Claim) error
	CreateCredential(userID string, publicKey string) (string, error)
}

type authRepository struct {
	spanner *gorm.DB
}

func NewRepository() AuthRepository {
	return &authRepository{
		spanner: utils.GetSpanner(),
	}
}
