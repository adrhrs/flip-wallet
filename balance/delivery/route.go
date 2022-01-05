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

	http.HandleFunc("/api/v1/balance/balance_topup", middleware.AuthUser(handler.incrementUserBalance))
	http.HandleFunc("/api/v1/balance/balance_read", middleware.AuthUser(handler.getUserBalanceInfo))
	http.HandleFunc("/api/v1/balance/transfer", middleware.AuthUser(handler.transferBalance))
	http.HandleFunc("/api/v1/balance/top_transactions_per_user", middleware.AuthUser(handler.getTopTransactionByUser))
	http.HandleFunc("/api/v1/balance/top_users", middleware.AuthUser(handler.getTopUserByTransaction))
}
