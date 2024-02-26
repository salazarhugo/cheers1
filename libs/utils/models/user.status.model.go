package models

import (
	"time"
)

// UserStatus User model
type UserStatus struct {
	UserId       string `gorm:"primarykey;column:UserId"`
	UserStatusId string `gorm:"primarykey;column:UserStatusId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Text         string
	Type         string
	DrinkId      string
}
