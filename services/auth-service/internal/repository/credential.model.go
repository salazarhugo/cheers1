package repository

import (
	"encoding/binary"
	"github.com/go-webauthn/webauthn/webauthn"
	"time"
)

// Credential model
type Credential struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	PublicKey []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CredentialToDomain(
	userID int64,
	credential *webauthn.Credential,
) *Credential {
	id := binary.LittleEndian.Uint64(credential.ID)
	return &Credential{
		ID:        int64(id),
		UserId:    userID,
		PublicKey: credential.PublicKey,
	}
}
