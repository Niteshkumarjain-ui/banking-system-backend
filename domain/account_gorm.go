package domain

import "time"

type Accounts struct {
	ID          uint `gorm:"primaryKey"`
	UserID      int
	AccountType string
	Balance     float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
