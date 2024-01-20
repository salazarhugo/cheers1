package repository

import "time"

// Friendship model
type Friendship struct {
	UserId1   string `gorm:"primarykey"`
	UserId2   string `gorm:"primarykey"`
	CreatedAt time.Time
}
