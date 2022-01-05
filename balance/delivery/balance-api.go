package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flip-clean/models"
)

func (b *BalanceHandler) getUserBalanceInfo(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		ctx            = req.Context()

		balance int
	)

	defer func() {
		data := models.GetUserBalanceInfoResponse{
			Balance: balance,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	balance, errUsecase := b.Usecase.GetUserBalanceInfo(ctx)
	if errUsecase != nil {
		log.Println(errUsecase)
		responseStatus = http.StatusInternalServerError
		return
	}

}

func (b *BalanceHandler) incrementUserBalance(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		ctx            = req.Context()

		incrementUserBalancePayload models.IncrementUserBalancePayload
		msg                         = "success"
	)

	defer func() {
		data := models.IncrementUserBalanceResponse{
			Message: msg,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	errDecode := json.NewDecoder(req.Body).Decode(&incrementUserBalancePayload)
	if errDecode != nil {
		log.Println(errDecode)
		msg = errDecode.Error()
		responseStatus = http.StatusBadRequest
		return
	}

	if incrementUserBalancePayload.Amount > 10000000 || incrementUserBalancePayload.Amount < 0 {
		msg = "invalid topup amount"
		responseStatus = http.StatusBadRequest
		return
	}

	b.Usecase.IncrementUserBalance(ctx, incrementUserBalancePayload)
}

func (b *BalanceHandler) transferBalance(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusNoContent
		ctx            = req.Context()

		transferBalancePayload models.TransferBalancePayload
		message                = "transfer success"
	)

	defer func() {
		data := models.IncrementUserBalanceResponse{
			Message: message,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	errDecode := json.NewDecoder(req.Body).Decode(&transferBalancePayload)
	if errDecode != nil {
		log.Println(errDecode)
		responseStatus = http.StatusBadRequest
		return
	}

	errTransfer := b.Usecase.TransferBalance(ctx, transferBalancePayload)
	if errTransfer != nil {
		switch errTransfer {
		case models.ErrUserExist:
			responseStatus = http.StatusNotFound
		case models.ErrInsufficientBalance:
			responseStatus = http.StatusBadRequest
		default:
			responseStatus = http.StatusInternalServerError
			log.Println(errTransfer)
		}
		message = errTransfer.Error()
		return
	}
}

func (b *BalanceHandler) getTopTransactionByUser(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusOK
		ctx            = req.Context()

		topTransaction []models.UserTransaction
	)

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(topTransaction)
	}()

	topTransaction = b.Usecase.GetTopTransactionByUser(ctx)
}

func (b *BalanceHandler) getTopUserByTransaction(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusOK
		topUser        []models.UserDebitTransaction
	)

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(topUser)
	}()

	topUser = b.Usecase.GetTopUserByTransaction()
}
