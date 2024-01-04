package repository

import (
	"github.com/go-webauthn/webauthn/webauthn"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"time"
)

// Credential model
type Credential struct {
	ID              []byte `gorm:"primarykey"`
	UserId          int64
	PublicKey       []byte
	AttestationType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CredentialToDomain(
	userID int64,
	credential *webauthn.Credential,
) *Credential {
	return &Credential{
		ID:              credential.ID,
		UserId:          userID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
	}
}

func (c *Credential) ToWebAuthnCredential() *webauthn.Credential {
	return &webauthn.Credential{
		ID:              c.ID,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       nil,
		Flags:           webauthn.CredentialFlags{},
		Authenticator:   webauthn.Authenticator{},
	}
}

func ToCredentialsPb(credentials []webauthn.Credential) []*pb.Credential {
	res := make([]*pb.Credential, 0)
	for _, credential := range credentials {
		res = append(res, ToCredentialPb(&credential))
	}
	return res
}

func ToCredentialPb(c *webauthn.Credential) *pb.Credential {
	transports := make([]string, 0)
	for _, transport := range c.Transport {
		transports = append(transports, string(transport))
	}
	return &pb.Credential{
		Id:              c.ID,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       transports,
	}
}
