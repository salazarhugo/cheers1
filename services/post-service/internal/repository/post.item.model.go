package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/audio"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"time"
)

// PostItem Post item model
type PostItem struct {
	PostId         string `gorm:"primarykey;column:PostId"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Caption        string
	City           string
	Location       string
	Photos         ArrayString `gorm:"type:VARCHAR(255)"`
	AudioUrl       string
	AudioWaveform  ArrayInt `gorm:"type:integer[]"`
	Username       string
	Name           string
	Verified       bool
	Picture        string
	DrinkID        int64
	DrinkName      string
	DrinkIcon      string
	LikeCount      int64
	HasViewerLiked bool
	Medias         PostMediaArray `gorm:"type:integer[]"`
}

// ToPostPb Domain -> Protobuf
func (p PostItem) ToPostPb() *postpb.Post {
	medias := make([]*postpb.PostMedia, 0)
	for _, media := range p.Medias {
		medias = append(medias, media.ToPostMediaPb())
	}

	return &postpb.Post{
		Id:           p.PostId,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		Address:      "",
		Privacy:      0,
		PostMedia:    medias,
		LocationName: p.Location,
		Drink: &postpb.Drink{
			Id:   p.DrinkID,
			Name: p.DrinkName,
			Icon: p.DrinkIcon,
		},
		Audio: &audio.Audio{
			Id:       0,
			Url:      p.AudioUrl,
			Waveform: p.AudioWaveform,
			Duration: 0,
		},
		Drunkenness:           0,
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

// ToUserPb Domain -> Protobuf
func (p PostItem) ToUserPb() *pb.User {
	return &pb.User{
		Id:       p.UserID,
		Name:     p.Name,
		Username: p.Username,
		Verified: p.Verified,
		Picture:  p.Picture,
	}
}
