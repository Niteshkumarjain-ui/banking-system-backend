package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"errors"
	"time"
)

func CreateAccount(request domain.AccountRequest) (response domain.AccountResponse, err error) {

	createAccount := domain.Accounts{
		UserID:      request.UserID,
		AccountType: request.AccountType,
		Balance:     request.Balance,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = outbound.DatabaseDriver.Create(&createAccount).Error

	if err != nil {
		return
	}

	response.Status = "Account Succesfully created."
	response.AccountId = createAccount.ID
	return
}

func GetAccount(accountId int, claims domain.JwtValidate) (response domain.GetAccountResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.First(&account, accountId).Error
	if err != nil {
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	response.ID = accountId
	response.AccuntType = account.AccountType
	response.Balance = account.Balance
	response.UserID = account.UserID
	return
}
