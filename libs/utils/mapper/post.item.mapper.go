package mapper

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/audio"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

type PostItem models.PostItem

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
		Longitude:    p.Longitude,
		Latitude:     p.Latitude,
		Drink: &postpb.Drink{
			Id:   p.DrinkID.StringVal,
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
