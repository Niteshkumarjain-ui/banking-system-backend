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

func GetAllAccount() (response []domain.GetAccountResponse, err error) {
	var accountRows []domain.Accounts
	var responseRow domain.GetAccountResponse
	err = outbound.DatabaseDriver.Find(&accountRows).Error
	if err != nil {
		return
	}

	for _, row := range accountRows {
		responseRow = domain.GetAccountResponse{}
		responseRow.ID = int(row.ID)
		responseRow.UserID = row.UserID
		responseRow.AccuntType = row.AccountType
		responseRow.Balance = row.Balance
		response = append(response, responseRow)
	}

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

func UpdateAccount(request domain.UpdateAccountRequest, claims domain.JwtValidate) (response domain.AccountResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.Where("id = ?", request.ID).First(&account).Error
	if err != nil {
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	account.AccountType = request.AccountType
	err = outbound.DatabaseDriver.Save(&account).Error
	if err != nil {
		return
	}

	response.AccountId = uint(request.ID)
	response.Status = "Account Updated Succesfully!"
	return
}

func DeleteAccount(accountId int, claims domain.JwtValidate) (response domain.AccountResponse, err error) {
	var account domain.Accounts

	err = outbound.DatabaseDriver.Where("id = ?", accountId).First(&account).Error
	if err != nil {
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}
	err = outbound.DatabaseDriver.Where("id = ?", accountId).Delete(&account).Error
	if err != nil {
		return
	}

	response.AccountId = uint(accountId)
	response.Status = "Account is Successfully Deleted."
	return
}
