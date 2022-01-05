package memory

import (
	"github.com/flip-clean/models"
)

const (
	userBalanceKey          = "balance-%v"
	topTransactionByUserKey = "top-transaction-%v"
	userDebitKey            = "debit-%v"
)

func (b *balanceMemoryRepo) GetBalance(username string) (balance int, isFound bool) {
	var result interface{}
	result, isFound = b.balanceDataSynced.Load(generateKey(userBalanceKey, username))
	balance = result.(int)
	return
}

func (b *balanceMemoryRepo) GetTopTransactionByUser(username string) []models.UserTransaction {
	return b.topTransactionByUser[generateKey(topTransactionByUserKey, username)]
}

func (b *balanceMemoryRepo) GetUserDebitValue(username string) int {
	return b.userDebitValue[generateKey(userDebitKey, username)]
}

func (b *balanceMemoryRepo) GetTopUserByTransaction() []models.UserDebitTransaction {
	return b.topUserByTransaction
}
