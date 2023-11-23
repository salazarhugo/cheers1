package repository

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
)

func (a *authRepository) CreateModerator(userID string) error {
	return a.CreateClaim(userID, constants.MODERATOR)
}
