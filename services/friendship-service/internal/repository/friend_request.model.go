package repository

import "time"

// FriendRequest model
type FriendRequest struct {
	UserId1   string `gorm:"primarykey"`
	UserId2   string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
