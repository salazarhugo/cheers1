package models

import (
	"time"
)

// UserStatusItem model
type UserStatusItem struct {
	UserId       string `gorm:"primarykey;column:UserId"`
	UserStatusId string `gorm:"primarykey;column:UserStatusId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Text         string
	Type         string
	DrinkId      string
	DrinkName    string
	DrinkIcon    string
	Name         string
	Username     string
	Picture      string
}
