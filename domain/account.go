package domain

type AccountRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	AccountType string `json:"account_type" binding:"required"`
}

type AccountResponse struct {
	AccountId uint   `json:"account_id"`
	Status    string `json:"status"`
}

type GetAccountResponse struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Balance    float64 `json:"balance"`
	AccuntType string  `json:"account_type"`
}

type UpdateAccountRequest struct {
	ID          int    `json:"id"`
	AccountType string `json:"account_type" binding:"required"`
}
