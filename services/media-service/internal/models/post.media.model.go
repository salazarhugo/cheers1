package models

import (
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"time"
)

type PostMedia struct {
	PostId               string `gorm:"primarykey;column:PostId"`
	PostMediaId          string `gorm:"primarykey;column:PostMediaId"`
	Url                  string
	Ref                  string
	Type                 string
	AccessibilityCaption string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// ToPostMediaPb Domain -> Protobuf
func (p PostMedia) ToPostMediaPb() *postpb.PostMedia {
	return &postpb.PostMedia{
		AccessibilityCaption: p.AccessibilityCaption,
		MediaType:            toMediaType(p.Type),
		ImageVersions: []*postpb.ImageVersion{
			&postpb.ImageVersion{
				Url:    p.Url,
				Width:  0,
				Height: 0,
			},
		},
	}
}

func toMediaType(mediaType string) postpb.MediaType {
	switch mediaType {
	case "IMAGE":
		return postpb.MediaType_IMAGE
	case "VIDEO":
		return postpb.MediaType_VIDEO
	default:
		return postpb.MediaType_IMAGE
	}
}
