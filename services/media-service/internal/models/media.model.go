package models

import (
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
