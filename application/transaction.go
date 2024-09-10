package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"errors"
)

func DepositFunds(request domain.DepositFundsRequest, claims domain.JwtValidate) (respones domain.DepositFundsResponse, err error) {
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

func WithdrawlFunds(request domain.DepositFundsRequest, claims domain.JwtValidate) (respones domain.DepositFundsResponse, err error) {
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
