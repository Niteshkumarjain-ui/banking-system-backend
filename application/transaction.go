package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"errors"
)

func DepositFunds(request domain.DepositWithdrawlFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.First(&account, request.AccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	account.Balance += request.Amount

	err = outbound.DatabaseDriver.Save(&account).Error
	if err != nil {
		return
	}

	transactions := domain.Transactions{
		AccountID:       uint(request.AccountID),
		TransactionType: "deposit",
		Amount:          request.Amount,
	}

	err = outbound.DatabaseDriver.Create(&transactions).Error
	if err != nil {
		return
	}

	respones.Status = "Deposit Successfully."
	respones.CurrentBalance = account.Balance
	return
}

func WithdrawlFunds(request domain.DepositWithdrawlFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.First(&account, request.AccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	if account.Balance < request.Amount {
		err = errors.New("Insufficent Balance.")
		return
	}

	account.Balance -= request.Amount

	err = outbound.DatabaseDriver.Save(&account).Error
	if err != nil {
		return
	}

	transactions := domain.Transactions{
		AccountID:       uint(request.AccountID),
		TransactionType: "withdrawl",
		Amount:          request.Amount,
	}

	err = outbound.DatabaseDriver.Create(&transactions).Error
	if err != nil {
		return
	}

	respones.Status = "Withdrawl Successfully."
	respones.CurrentBalance = account.Balance
	return
}

func TransferFunds(request domain.TransferFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	var fromAccount, toAccount domain.Accounts

	err = outbound.DatabaseDriver.First(&fromAccount, request.FromAccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	err = outbound.DatabaseDriver.First(&toAccount, request.ToAccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(fromAccount.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	if fromAccount.Balance < request.Amount {
		err = errors.New("Insufficent Balance.")
		return
	}

	fromAccount.Balance -= request.Amount
	toAccount.Balance += request.Amount

	err = outbound.DatabaseDriver.Save(&fromAccount).Error
	if err != nil {
		return
	}
	err = outbound.DatabaseDriver.Save(&toAccount).Error
	if err != nil {
		return
	}

	fromTransactions := domain.Transactions{
		AccountID:       uint(request.FromAccountID),
		TransactionType: "transfer_out",
		Amount:          request.Amount,
	}
	toTransactions := domain.Transactions{
		AccountID:       uint(request.ToAccountID),
		TransactionType: "transfer_in",
		Amount:          request.Amount,
	}

	err = outbound.DatabaseDriver.Create(&fromTransactions).Error
	if err != nil {
		return
	}
	err = outbound.DatabaseDriver.Create(&toTransactions).Error
	if err != nil {
		return
	}

	respones.Status = "transfer fund Successfully."
	respones.CurrentBalance = fromAccount.Balance
	return
}
