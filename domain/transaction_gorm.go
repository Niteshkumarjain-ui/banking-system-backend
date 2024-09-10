package domain

import "time"

type Transactions struct {
	ID              uint `gorm:"primaryKey"`
	AccountID       uint
	TransactionType string
	Amount          float64
	CreatedAt       time.Time
}
