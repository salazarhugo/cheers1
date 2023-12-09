package repository

import (
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"strconv"
	"time"
)

// Post model
type Post struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	DrinkID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
	Location  string
}

// PostWithUserInfo Post item model
type PostWithUserInfo struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
	Location  string
	Username  string
	Name      string
	Verified  bool
	Picture   string
	DrinkID   int64
	DrinkName string
	DrinkIcon string
}

func ToPost(post *postpb.Post) *Post {
	return &Post{
		ID:        post.Id,
		UserID:    post.CreatorId,
		CreatedAt: time.Unix(post.CreateTime, 0),
		UpdatedAt: time.Time{},
		Caption:   post.Caption,
		City:      "",
		Location:  post.LocationName,
		DrinkID:   1,
	}
}

func (p Post) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:           p.ID,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		Photos:       []string{},
		LocationName: p.Location,
		CreateTime:   p.CreatedAt.Unix(),
		Drink:        &postpb.Drink{},
	}
}

func (p PostWithUserInfo) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:           p.ID,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		Address:      "",
		Privacy:      0,
		Photos:       []string{p.Picture},
		LocationName: p.Location,
		Drink: &postpb.Drink{
			Id:   strconv.FormatInt(p.DrinkID, 10),
			Name: p.DrinkName,
			Icon: p.DrinkIcon,
		},
		Drunkenness:           0,
		Type:                  0,
		CreateTime:            p.CreatedAt.Unix(),
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
