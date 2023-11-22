package repository

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/google/uuid"
)

func (a *authRepository) GetOrCreateUser(
	username string,
) (*User, error) {
	user, err := a.GetUserByUsername(username)
	if user != nil {
		return user, err
	}

	newUser := &User{
		ID:       uuid.NewString(),
		Username: username,
		Verified: false,
		AuthnId:  int64(randomUint64()),
	}

	// Insert user into database
	err = a.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	// Get user from database
	user, err = a.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
