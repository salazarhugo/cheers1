package repository

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
)

func (a authRepository) CreateAdmin(userID string) error {
	return a.CreateClaim(userID, constants.ADMIN)
}
