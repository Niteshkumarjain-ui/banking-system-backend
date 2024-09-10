package domain

import "time"

type DepositWithdrawlFundsRequest struct {
	AccountID int     `json:"account_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type TransactionResponse struct {
	Status         string  `json:"status"`
	CurrentBalance float64 `josn:"current_balance"`
}

type TransferFundsRequest struct {
	FromAccountID int     `json:"from_account_id" binding:"required"`
	ToAccountID   int     `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

type GetAccountStatement struct {
	ID              uint      `json:"id"`
	AccountID       int       `json:"account_id"`
	TransactionType string    `json:"transcation_type"`
	Amount          float64   `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}
