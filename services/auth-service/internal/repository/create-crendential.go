package repository

import "github.com/google/uuid"

func (a authRepository) CreateCredential(
	userID string,
	publicKey string,
) (string, error) {
	db := a.spanner

	credential := &Credential{
		ID:        uuid.NewString(),
		UserId:    userID,
		PublicKey: publicKey,
	}

	result := db.Create(&credential)
	if result.Error != nil {
		return "", result.Error
	}

	return credential.ID, nil
}
