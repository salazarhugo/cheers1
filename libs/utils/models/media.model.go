package models

import (
	"github.com/google/uuid"
	"time"
)

type Media struct {
	MediaId              string `gorm:"primarykey;column:MediaId"`
	Url                  string
	Ref                  string
	Type                 string
	AccessibilityCaption string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (m Media) ToPostMedia(postID string) *PostMedia {
	return &PostMedia{
		PostId:               postID,
		PostMediaId:          uuid.NewString(),
		Url:                  m.Url,
		Ref:                  m.Ref,
		Type:                 m.Type,
		AccessibilityCaption: m.AccessibilityCaption,
	}
}
