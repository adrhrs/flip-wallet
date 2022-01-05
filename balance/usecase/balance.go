package usecase

import (
	"context"
	"math"

	"github.com/flip-clean/models"
	"github.com/flip-clean/pkg"
)

func (b *balanceUsecase) IncrementUserBalance(ctx context.Context, param models.IncrementUserBalancePayload) (err error) {
	username := pkg.GetUsernameFromCtx(ctx)
	err = b.BalanceRepo.IncBalance(username, param.Amount)
	return
}

func (b *balanceUsecase) GetUserBalanceInfo(ctx context.Context) (balance int, err error) {
	username := pkg.GetUsernameFromCtx(ctx)
	balance, isFound := b.BalanceRepo.GetBalance(username)
	if !isFound {
		err = models.ErrUserNotFound
		return
	}
	return
}

func (b *balanceUsecase) TransferBalance(ctx context.Context, param models.TransferBalancePayload) (err error) {
	username := pkg.GetUsernameFromCtx(ctx)
	balance, isFound := b.BalanceRepo.GetBalance(username)
	if !isFound {
		err = models.ErrUserNotFound
		return
	}
	if balance < param.Amount {
		err = models.ErrInsufficientBalance
		return
	}
	isUserDestinationExist := b.UserUsecase.IsUserExist(param.ToUsername)
	if !isUserDestinationExist {
		err = models.ErrUserNotFound
		return
	}

	err = b.startTransferOperation(username, param)

	return
}

func (b *balanceUsecase) startTransferOperation(username string, param models.TransferBalancePayload) (err error) {
	var (
		errCode                  int
		prevTotalDebitSender     int
		prevTopTransactionSender []models.UserTransaction
		prevTopTransactionTarget []models.UserTransaction
	)
	defer func() {
		// if fail dec sender balance no need to rollback since this is the first activity
		if errCode > 1 {
			b.rollBackOperations(errCode, username, param, prevTopTransactionSender, prevTopTransactionTarget, prevTotalDebitSender)
		}
	}()

	err = b.BalanceRepo.IncBalance(username, -param.Amount)
	if err != nil {
		err = models.ErrDecSenderBalance
		errCode = 1
		return
	}
	err = b.BalanceRepo.IncBalance(param.ToUsername, param.Amount)
	if err != nil {
		err = models.ErrIncTargetBalance
		errCode = 2
		return
	}

	senderTopTransactionData := b.BalanceRepo.GetTopTransactionByUser(username)
	prevTopTransactionSender = senderTopTransactionData
	senderTopTransactionData = evaluateTopTransaction(senderTopTransactionData, -param.Amount, username)
	err = b.BalanceRepo.SetTopTransactionByUser(username, senderTopTransactionData)
	if err != nil {
		err = models.ErrSenderSetTopTransactionByUser
		errCode = 3
		return
	}

	targetTopTransactionData := b.BalanceRepo.GetTopTransactionByUser(param.ToUsername)
	prevTopTransactionTarget = targetTopTransactionData
	targetTopTransactionData = evaluateTopTransaction(targetTopTransactionData, param.Amount, param.ToUsername)
	err = b.BalanceRepo.SetTopTransactionByUser(param.ToUsername, targetTopTransactionData)
	if err != nil {
		err = models.ErrTargetSetTopTransactionByUser
		errCode = 4
		return
	}

	err = b.BalanceRepo.SetUserDebitValue(username, param.Amount)
	if err != nil {
		err = models.ErrSetTotalDebitSender
		errCode = 5
		return
	}
	senderTotalDebit := b.BalanceRepo.GetUserDebitValue(username)
	prevTotalDebitSender = senderTotalDebit - param.Amount
	topDebitur := b.BalanceRepo.GetTopUserByTransaction()
	topDebitur = evaluateTopUserByTransaction(topDebitur, senderTotalDebit, username)
	err = b.BalanceRepo.SetTopUserByTransaction(topDebitur)
	if err != nil {
		err = models.ErrSetTopUserByTransaction
		errCode = 6
		return
	}

	return
}

func (b *balanceUsecase) rollBackOperations(
	errCode int, username string,
	param models.TransferBalancePayload,
	prevTopTransactionSender []models.UserTransaction,
	prevTopTransactionTarget []models.UserTransaction,
	prevTotalDebitSender int,
) {
	for i := 2; i <= errCode; i++ {
		switch i {
		case 2:
			// re-increase sender balance
			b.BalanceRepo.IncBalance(username, param.Amount)
		case 3:
			// re-decrease target balance
			b.BalanceRepo.IncBalance(param.ToUsername, -param.Amount)
		case 4:
			// re-set top transaction sender
			b.BalanceRepo.SetTopTransactionByUser(username, prevTopTransactionSender)
		case 5:
			// re-set top transaction target
			b.BalanceRepo.SetTopTransactionByUser(param.ToUsername, prevTopTransactionTarget)
		case 6:
			// re-set total debit
			b.BalanceRepo.SetUserDebitValue(username, prevTotalDebitSender)
		}
	}
}

func evaluateTopUserByTransaction(topDebitur []models.UserDebitTransaction, amount int, username string) []models.UserDebitTransaction {
	if len(topDebitur) > 0 {
		var absLeastAmountOfTransaction float64
		absCurrentTrxValue := math.Abs(float64(amount))
		if len(topDebitur) == 10 {
			absLeastAmountOfTransaction = math.Abs(float64(topDebitur[len(topDebitur)-1].TransactedValue))
		}
		if absCurrentTrxValue > absLeastAmountOfTransaction {
			for id, trx := range topDebitur {
				absTopUserTrxValue := math.Abs(float64(trx.TransactedValue))
				if absCurrentTrxValue > absTopUserTrxValue {
					topDebitur = insertTopUserArr(topDebitur, id, models.UserDebitTransaction{
						Username:        username,
						TransactedValue: amount,
					})
					break
				} else if id == len(topDebitur)-1 {
					topDebitur = append(topDebitur, models.UserDebitTransaction{
						Username:        username,
						TransactedValue: amount,
					})
				}
			}
		}
		if len(topDebitur) > 10 {
			topDebitur = topDebitur[:10]
		}
	} else {
		topDebitur = append(topDebitur, models.UserDebitTransaction{
			Username:        username,
			TransactedValue: amount,
		})
	}

	return topDebitur
}

func insertTopUserArr(data []models.UserDebitTransaction, i int, trx models.UserDebitTransaction) []models.UserDebitTransaction {
	if i == len(data) {
		return append(data, trx)
	}
	data = append(data[:i+1], data[i:]...)
	data[i] = trx
	return data
}

func evaluateTopTransaction(topUserTransaction []models.UserTransaction, amount int, username string) []models.UserTransaction {
	if len(topUserTransaction) > 0 {
		var absLeastAmountOfTransaction float64
		absCurrentTrxValue := math.Abs(float64(amount))
		if len(topUserTransaction) == 10 {
			absLeastAmountOfTransaction = math.Abs(float64(topUserTransaction[len(topUserTransaction)-1].Amount))
		}
		if absCurrentTrxValue > absLeastAmountOfTransaction {
			for id, trx := range topUserTransaction {
				absTopUserTrxValue := math.Abs(float64(trx.Amount))
				if absCurrentTrxValue > absTopUserTrxValue {
					topUserTransaction = insertTopTransactionArr(topUserTransaction, id, models.UserTransaction{
						Username: username,
						Amount:   amount,
					})
					break
				} else if id == len(topUserTransaction)-1 {
					topUserTransaction = append(topUserTransaction, models.UserTransaction{
						Username: username,
						Amount:   amount,
					})
				}
			}
		}
		if len(topUserTransaction) > 10 {
			topUserTransaction = topUserTransaction[:10]
		}
	} else {
		topUserTransaction = append(topUserTransaction, models.UserTransaction{
			Username: username,
			Amount:   amount,
		})
	}

	return topUserTransaction
}

func insertTopTransactionArr(data []models.UserTransaction, i int, trx models.UserTransaction) []models.UserTransaction {
	if i == len(data) {
		return append(data, trx)
	}
	data = append(data[:i+1], data[i:]...)
	data[i] = trx
	return data
}

func (b *balanceUsecase) GetTopTransactionByUser(ctx context.Context) []models.UserTransaction {
	username := pkg.GetUsernameFromCtx(ctx)
	return b.BalanceRepo.GetTopTransactionByUser(username)
}

func (b *balanceUsecase) GetTopUserByTransaction() []models.UserDebitTransaction {
	return b.BalanceRepo.GetTopUserByTransaction()
}
