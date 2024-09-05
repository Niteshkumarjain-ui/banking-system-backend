package domain

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
