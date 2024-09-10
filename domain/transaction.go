package domain

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
