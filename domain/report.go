package domain

import "time"

type GetAccountBalanceResponse struct {
	ID               int     `json:"id"`
	AvailableBalance float64 `json:"available_balance"`
}

type GetDailyTransactionReportResponse struct {
	TransactionType string    `json:"transcation_type"`
	Amount          float64   `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}

type GetFinancialReportResponse struct {
	TotalDeposit     float64 `json:"total_deposit"`
	TotalWithdrawl   float64 `json:"total_withdrawl"`
	AvailableBalance float64 `json:"available_balance"`
	TotalTransferIn  float64 `json:"total_transfer_in"`
	TotalTransferOut float64 `json:"total_transfer_out"`
}
