package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"banking-system-backend/util"
	"context"
	"errors"
)

func DepositFunds(ctx context.Context, request domain.DepositWithdrawlFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	_, span := util.Tracer.Start(ctx, "DepositFunds")
	defer span.End()
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

func WithdrawlFunds(ctx context.Context, request domain.DepositWithdrawlFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	_, span := util.Tracer.Start(ctx, "WithdrawlFunds")
	defer span.End()
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

func TransferFunds(ctx context.Context, request domain.TransferFundsRequest, claims domain.JwtValidate) (respones domain.TransactionResponse, err error) {
	_, span := util.Tracer.Start(ctx, "TransferFunds")
	defer span.End()
	var fromAccount, toAccount domain.Accounts

	err = outbound.DatabaseDriver.First(&fromAccount, request.FromAccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(fromAccount.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	err = outbound.DatabaseDriver.First(&toAccount, request.ToAccountID).Error
	if err != nil {
		err = errors.New("Account Not Found.")
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

func GetAccountStatement(ctx context.Context, accountId int, claims domain.JwtValidate) (response []domain.GetAccountStatement, err error) {
	_, span := util.Tracer.Start(ctx, "GetAccountStatement")
	defer span.End()
	var transcationsRows []domain.Transactions
	var responseRow domain.GetAccountStatement
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

	err = outbound.DatabaseDriver.Where("account_id = ?", accountId).Find(&transcationsRows).Error
	if err != nil {
		return
	}

	for _, row := range transcationsRows {
		responseRow = domain.GetAccountStatement{}
		responseRow.ID = row.ID
		responseRow.Amount = row.Amount
		responseRow.TransactionType = row.TransactionType
		responseRow.CreatedAt = row.CreatedAt
		responseRow.AccountID = int(row.AccountID)
		response = append(response, responseRow)
	}

	return
}

func GetTransaction(ctx context.Context, transactionId int, claims domain.JwtValidate) (response domain.GetAccountStatement, err error) {
	_, span := util.Tracer.Start(ctx, "GetTransaction")
	defer span.End()
	var account domain.Accounts
	var transaction domain.Transactions

	err = outbound.DatabaseDriver.First(&transaction, transactionId).Error
	if err != nil {
		err = errors.New("Transcation Not Found.")
		return
	}

	err = outbound.DatabaseDriver.First(&account, transaction.AccountID).Error
	if err != nil {
		err = errors.New("Account Not Found")
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(account.UserID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this account")
		return
	}

	response.ID = transaction.ID
	response.Amount = transaction.Amount
	response.TransactionType = transaction.TransactionType
	response.CreatedAt = transaction.CreatedAt
	response.AccountID = int(transaction.AccountID)
	return
}
