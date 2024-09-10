package domain

type DepositFundsRequest struct {
	AccountID int     `json:"account_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type DepositFundsResponse struct {
	Status         string  `json:"status"`
	CurrentBalance float64 `josn:"current_balance"`
}
