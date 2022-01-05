package usecase

import (
	"context"

	repo "github.com/flip-clean/balance/repository"
	"github.com/flip-clean/models"
	userUsecase "github.com/flip-clean/user/usecase"
)

type balanceUsecase struct {
	BalanceRepo repo.BalanceRepo
	UserUsecase userUsecase.UserUsecase
}

// NewBalanceUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewBalanceUsecase(balanceRepo repo.BalanceRepo, userUsecase userUsecase.UserUsecase) BalanceUsecase {
	return &balanceUsecase{
		BalanceRepo: balanceRepo,
		UserUsecase: userUsecase,
	}
}

// BalanceUsecase represent the BalanceUsecase's usecases
type BalanceUsecase interface {
	IncrementUserBalance(ctx context.Context, param models.IncrementUserBalancePayload) (err error)
	GetUserBalanceInfo(ctx context.Context) (balance int, err error)
	TransferBalance(ctx context.Context, param models.TransferBalancePayload) (err error)
	GetTopTransactionByUser(ctx context.Context) []models.UserTransaction
	GetTopUserByTransaction() []models.UserDebitTransaction
}
