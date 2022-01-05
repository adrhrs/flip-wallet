package repository

import "github.com/flip-clean/models"

//go:generate moq -out repo_mock.go . BalanceRepo
type BalanceRepo interface {
	IncBalance(username string, increment int) (err error)
	GetBalance(username string) (balance int, isFound bool)
	GetTopTransactionByUser(username string) []models.UserTransaction
	SetTopTransactionByUser(username string, value []models.UserTransaction) (err error)
	GetUserDebitValue(username string) int
	SetUserDebitValue(username string, value int) (err error)
	GetTopUserByTransaction() []models.UserDebitTransaction
	SetTopUserByTransaction(value []models.UserDebitTransaction) (err error)
}
