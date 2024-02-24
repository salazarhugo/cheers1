package model

import (
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// UserWithViewer model
type UserWithViewer struct {
	UserId             string `gorm:"primarykey;column:UserId"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Username           string
	Name               string
	Picture            string
	Verified           bool
	Banner             string
	Friend             bool
	Requested          bool
	Website            string
	Bio                string
	Email              string
	HasRequestedViewer bool
	PostCount          int32
	FriendCount        int32
}

func (u UserWithViewer) ToUserPb() *userpb.User {
	return &userpb.User{
		Id:                 u.UserId,
		Name:               u.Name,
		Email:              u.Email,
		Verified:           u.Verified,
		Username:           u.Username,
		Picture:            u.Picture,
		Bio:                u.Bio,
		Website:            u.Website,
		Birthday:           "",
		Gender:             0,
		PhoneNumber:        "",
		CreateTime:         0,
		RegistrationTokens: nil,
		IsBusinessAccount:  false,
		IsAdmin:            false,
		IsModerator:        false,
		Banner:             u.Banner,
	}
}

func (u UserWithViewer) ToUserItemPb() *userpb.UserItem {
	return &userpb.UserItem{
		Id:                 u.UserId,
		Name:               u.Name,
		Username:           u.Username,
		Verified:           u.Verified,
		Picture:            u.Picture,
		Banner:             u.Banner,
		HasFollowed:        false,
		StoryState:         0,
		Friend:             u.Friend,
		Requested:          u.Requested,
		HasRequestedViewer: u.HasRequestedViewer,
	}
}

func ToUserItemsPb2(users []*UserWithViewer) []*userpb.UserItem {
	items := make([]*userpb.UserItem, 0)
	for _, u := range users {
		items = append(items, u.ToUserItemPb())
	}
	return items
}
