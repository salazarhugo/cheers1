package repository

import (
	"fmt"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/app"
	"sync"
)

type userdb struct {
	users map[string]*app.AuthnUser
	mu    sync.RWMutex
}

var db *userdb

// DB returns a userdb singleton
func DB() *userdb {

	if db == nil {
		db = &userdb{
			users: make(map[string]*app.AuthnUser),
		}
	}

	return db
}

// GetAuthnUser returns a *AuthnUser by the user's username
func (db *userdb) GetAuthnUser(name string) (*app.AuthnUser, error) {

	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[name]
	if !ok {
		return &app.AuthnUser{}, fmt.Errorf("error getting user '%s': does not exist", name)
	}

	return user, nil
}

// PutAuthnUser stores a new user by the user's username
func (db *userdb) PutAuthnUser(user *app.AuthnUser) {

	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.WebAuthnName()] = user
}
