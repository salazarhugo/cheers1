package repository

import (
	"cloud.google.com/go/spanner"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/audio"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"time"
)

// Post model
type Post struct {
	PostId        string `gorm:"primarykey;column:PostId"`
	UserID        string
	DrinkID       spanner.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Caption       string
	City          string
	Location      string
	Latitude      float64
	Longitude     float64
	AudioUrl      string
	AudioWaveform ArrayInt `gorm:"type:integer[]"`
}

// ToPostPb Domain -> Protobuf
func (p Post) ToPostPb() *postpb.Post {
	return &postpb.Post{
		Id:           p.PostId,
		CreatorId:    p.UserID,
		Caption:      p.Caption,
		PostMedia:    nil,
		LocationName: p.Location,
		Longitude:    p.Longitude,
		Latitude:     p.Latitude,
		CreateTime:   p.CreatedAt.Unix(),
		Drink: &postpb.Drink{
			Id: p.DrinkID.StringVal,
		},
		Audio: &audio.Audio{
			Id:       0,
			Url:      p.AudioUrl,
			Waveform: p.AudioWaveform,
			Duration: 0,
		},
	}
}

func ToStringArray(s []spanner.NullString) []string {
	res := make([]string, 0)
	for _, str := range s {
		res = append(res, str.StringVal)
	}
	return res
}
