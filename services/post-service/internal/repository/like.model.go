package repository

import (
	"time"
)

// Like model
type Like struct {
	UserID    string `gorm:"primarykey"`
	PostID    string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// LikeWithUser model
type LikeWithUser struct {
	UserID    string `gorm:"primarykey"`
	PostID    string `gorm:"primarykey"`
	Username  string
	Name      string
	Verified  bool
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
