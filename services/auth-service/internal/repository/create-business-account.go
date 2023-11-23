package repository

import "github.com/salazarhugo/cheers1/services/auth-service/internal/constants"

func (a *authRepository) CreateBusinessAccount(userID string) error {
	return a.CreateClaim(userID, constants.BUSINESS)
}
