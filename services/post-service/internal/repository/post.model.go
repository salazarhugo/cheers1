package repository

import (
	"cloud.google.com/go/spanner"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// Post model
type Post struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	DrinkID   spanner.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
	Location  string
	Photos    ArrayString `gorm:"type:VARCHAR(255)"`
}

// PostWithUserInfo Post item model
type PostWithUserInfo struct {
	ID             string `gorm:"primarykey"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Caption        string
	City           string
	Location       string
	Photos         ArrayString `gorm:"type:VARCHAR(255)"`
	Username       string
	Name           string
	Verified       bool
	Picture        string
	DrinkID        int64
	DrinkName      string
	DrinkIcon      string
	Likes          int64
	HasViewerLiked bool
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
		DrinkID:   spanner.NullInt64{Int64: post.Drink.Id},
		Photos:    post.Photos,
	}
}

func (p Post) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:           p.ID,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		Photos:       p.Photos,
		LocationName: p.Location,
		CreateTime:   p.CreatedAt.Unix(),
		Drink: &postpb.Drink{
			Id: p.DrinkID.Int64,
		},
	}
}

func ToSpannerStringArray(s []string) []spanner.NullString {
	res := make([]spanner.NullString, 0)
	for _, str := range s {
		res = append(res, spanner.NullString{StringVal: str, Valid: true})
	}
	return res
}

func ToStringArray(s []spanner.NullString) []string {
	res := make([]string, 0)
	for _, str := range s {
		res = append(res, str.StringVal)
	}
	return res
}

func (p PostWithUserInfo) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:           p.ID,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		Address:      "",
		Privacy:      0,
		Photos:       p.Photos,
		LocationName: p.Location,
		Drink: &postpb.Drink{
			Id:   p.DrinkID,
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
