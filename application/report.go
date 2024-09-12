package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"banking-system-backend/util"
	"errors"
	"time"
)

func GetAccountBalance(accountId int, claims domain.JwtValidate) (response domain.GetAccountBalanceResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.First(&account, accountId).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	response.ID = accountId
	response.AvailableBalance = account.Balance
	return
}

func GetFinancialReport(accountId int, claims domain.JwtValidate) (response domain.GetFinancialReportResponse, err error) {
	var account domain.Accounts
	var transactions []domain.Transactions
	err = outbound.DatabaseDriver.First(&account, accountId).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	startFinancialYear, endFinancialYear := util.GetFinancialYear()

	err = outbound.DatabaseDriver.Where("account_id = ? AND created_at BETWEEN ? AND ?", accountId, startFinancialYear, endFinancialYear).Find(&transactions).Error

	for _, transaction := range transactions {
		if transaction.TransactionType == "deposit" {
			response.TotalDeposit += transaction.Amount
		} else if transaction.TransactionType == "withdrawl" {
			response.TotalWithdrawl += transaction.Amount
		} else if transaction.TransactionType == "transfer_in" {
			response.TotalTransferIn += transaction.Amount
		} else if transaction.TransactionType == "transfer_out" {
			response.TotalTransferOut += response.TotalTransferOut
		}
	}
	response.AvailableBalance = account.Balance
	return
}

func GetDailyTransactionReport() (response []domain.GetDailyTransactionReportResponse, err error) {
	var transcationsRows []domain.Transactions
	var responseRow domain.GetDailyTransactionReportResponse

	today := time.Now().UTC().Truncate(24 * time.Hour)
	err = outbound.DatabaseDriver.Where("created_At >= ?", today).Find(&transcationsRows).Error
	if err != nil {
		return
	}

	for _, row := range transcationsRows {
		responseRow = domain.GetDailyTransactionReportResponse{}
		responseRow.Amount = row.Amount
		responseRow.TransactionType = row.TransactionType
		responseRow.CreatedAt = row.CreatedAt
		response = append(response, responseRow)
	}
	return
}
