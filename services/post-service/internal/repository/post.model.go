package repository

import (
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// Post model
type Post struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
}

// PostWithUserInfo Post item model
type PostWithUserInfo struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
	Username  string
	Name      string
	Verified  bool
	Picture   string
}

func (p PostWithUserInfo) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:                    p.ID,
		CreatorId:             p.UserID,
		Caption:               p.Caption,
		Address:               "",
		Privacy:               0,
		Photos:                []string{p.Picture},
		LocationName:          p.City,
		Drink:                 "",
		Drunkenness:           0,
		Type:                  0,
		CreateTime:            0,
		CanComment:            false,
		CanShare:              false,
		Ratio:                 0,
		Latitude:              0,
		Longitude:             0,
		LastCommentText:       "",
		LastCommentUsername:   "",
		LastCommentCreateTime: 0,
	}
}

func (p PostWithUserInfo) ToUserPb() *pb.User {
	return &pb.User{
		Id:       p.UserID,
		Name:     p.Name,
		Username: p.Username,
		Verified: p.Verified,
		Picture:  p.Picture,
	}
}
