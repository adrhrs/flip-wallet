package delivery

import (
	"net/http"

	balanceUsecase "github.com/flip-clean/balance/usecase"
	"github.com/flip-clean/middleware"
)

type BalanceHandler struct {
	Usecase balanceUsecase.BalanceUsecase
}

func NewBalanceHandler(usecase balanceUsecase.BalanceUsecase) {
	handler := &BalanceHandler{
		Usecase: usecase,
	}

	http.HandleFunc("/balance_topup", middleware.AuthUser(handler.incrementUserBalance))
	http.HandleFunc("/balance_read", middleware.AuthUser(handler.getUserBalanceInfo))
	http.HandleFunc("/transfer", middleware.AuthUser(handler.transferBalance))
	http.HandleFunc("/top_transactions_per_user", middleware.AuthUser(handler.getTopTransactionByUser))
	http.HandleFunc("/top_users", middleware.AuthUser(handler.getTopUserByTransaction))
}
