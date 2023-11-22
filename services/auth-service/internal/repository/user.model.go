package repository

import (
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// User model
type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Name      string
	Picture   string
	Verified  bool
	AuthnId   int64
}

func (u User) ToUserPb() *userpb.User {
	return &userpb.User{
		Id:                 u.ID,
		Name:               u.Name,
		Email:              u.Email,
		Verified:           u.Verified,
		Username:           u.Username,
		Picture:            u.Picture,
		Bio:                "",
		Website:            "",
		Birthday:           "",
		Gender:             0,
		PhoneNumber:        "",
		CreateTime:         u.CreatedAt.Unix(),
		RegistrationTokens: nil,
		IsBusinessAccount:  false,
		IsAdmin:            false,
		IsModerator:        false,
		Banner:             "",
	}
}

func (u User) ToAuthnUser() *AuthnUser {
	return NewUser(uint64(u.AuthnId), u.Username, u.Username)
}
