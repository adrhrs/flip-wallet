package memory

import (
	"sync"

	repo "github.com/flip-clean/balance/repository"
	"github.com/flip-clean/models"
)

type balanceMemoryRepo struct {
	topTransactionByUser map[string][]models.UserTransaction
	userDebitValue       map[string]int
	topUserByTransaction []models.UserDebitTransaction

	balanceDataSynced sync.Map
}

func NewBalanceMemoryRepo(
	topTransactionByUser map[string][]models.UserTransaction,
	userDebitValue map[string]int,
	topUserByTransaction []models.UserDebitTransaction,
	balanceDataSynced sync.Map,
) repo.BalanceRepo {
	return &balanceMemoryRepo{
		topTransactionByUser: topTransactionByUser,
		userDebitValue:       userDebitValue,
		topUserByTransaction: topUserByTransaction,
		balanceDataSynced:    balanceDataSynced,
	}
}
