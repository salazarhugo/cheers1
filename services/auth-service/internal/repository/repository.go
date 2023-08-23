package repository

import "github.com/salazarhugo/cheers1/services/auth-service/internal/constants"

type AuthRepository interface {
	CreateAdmin(userID string) error
	CreateModerator(userID string) error
	CreateBusinessAccount(userID string) error
	VerifyUser(userID string) error
	CreateClaim(userID string, claim constants.Claim) error
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}
