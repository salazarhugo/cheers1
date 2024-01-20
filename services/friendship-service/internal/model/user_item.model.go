package model

import (
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// UserItem model
type UserItem struct {
	ID                 string `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Username           string
	Name               string
	Picture            string
	Verified           bool
	Banner             string
	Friend             bool
	Requested          bool
	HasRequestedViewer bool
}

func (u UserItem) ToUserItemPb() *userpb.UserItem {
	return &userpb.UserItem{
		Id:                 u.ID,
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

func ToUserItemsPb2(users []*UserItem) []*userpb.UserItem {
	items := make([]*userpb.UserItem, 0)
	for _, u := range users {
		items = append(items, u.ToUserItemPb())
	}
	return items
}
