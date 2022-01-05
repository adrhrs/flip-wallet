package memory

import (
	"fmt"

	"github.com/flip-clean/models"
)

func generateKey(key, username string) string {
	return fmt.Sprintf(key, username)
}

func (b *balanceMemoryRepo) SetTopUserByTransaction(value []models.UserDebitTransaction) (err error) {
	b.topUserByTransaction = value
	return
}

func (b *balanceMemoryRepo) SetUserDebitValue(username string, value int) (err error) {
	b.userDebitValue[generateKey(userDebitKey, username)] += value
	return
}

func (b *balanceMemoryRepo) SetTopTransactionByUser(username string, value []models.UserTransaction) (err error) {
	b.topTransactionByUser[generateKey(topTransactionByUserKey, username)] = value
	return
}

func (b *balanceMemoryRepo) IncBalance(username string, increment int) (err error) {
	b.balanceDataSynced.Store(generateKey(userBalanceKey, username), increment)
	return
}
