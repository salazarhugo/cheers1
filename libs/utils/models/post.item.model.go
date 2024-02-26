package models

import (
	"cloud.google.com/go/spanner"
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
	Latitude       float64
	Longitude      float64
	Photos         ArrayString `gorm:"type:VARCHAR(255)"`
	AudioUrl       string
	AudioWaveform  ArrayInt `gorm:"type:integer[]"`
	Username       string
	Name           string
	Verified       bool
	Picture        string
	DrinkID        spanner.NullString
	DrinkName      string
	DrinkIcon      string
	LikeCount      int64
	HasViewerLiked bool
	Medias         PostMediaArray `gorm:"type:integer[]"`
}
