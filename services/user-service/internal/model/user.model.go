package model

import (
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// User model
type User struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Username    string
	Email       string
	Name        string
	Picture     string
	Verified    bool
	AuthnId     int64
	Bio         string
	Website     string
	Banner      string
	PhoneNumber string
	IsAdmin     bool
}

func (u User) ToUserPb() *userpb.User {
	return &userpb.User{
		Id:                 u.ID,
		Name:               u.Name,
		Email:              u.Email,
		Verified:           u.Verified,
		Username:           u.Username,
		Picture:            u.Picture,
		Bio:                u.Bio,
		Website:            u.Website,
		Birthday:           "",
		Gender:             0,
		PhoneNumber:        u.PhoneNumber,
		CreateTime:         u.CreatedAt.Unix(),
		RegistrationTokens: nil,
		IsBusinessAccount:  false,
		IsAdmin:            u.IsAdmin,
		IsModerator:        false,
		Banner:             u.Banner,
	}
}

func (u User) ToUserItemPb() *userpb.UserItem {
	return &userpb.UserItem{
		Id:                 u.ID,
		Name:               u.Name,
		Username:           u.Username,
		Verified:           u.Verified,
		Picture:            u.Picture,
		HasFollowed:        false,
		StoryState:         0,
		Friend:             false,
		Requested:          false,
		HasRequestedViewer: false,
		Banner:             u.Banner,
	}
}

func ToUserItemsPb(users []*User) []*userpb.UserItem {
	items := make([]*userpb.UserItem, 0)
	for _, u := range users {
		items = append(items, u.ToUserItemPb())
	}
	return items
}
