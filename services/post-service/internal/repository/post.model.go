package repository

import (
	"time"
)

// Post model
type Post struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Caption   string
	City      string
}
