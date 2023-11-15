package repository

import "time"

// Credential model
type Credential struct {
	ID        string `gorm:"primarykey"`
	UserId    string
	PublicKey string
	CreatedAt time.Time
	UpdatedAt time.Time
}
