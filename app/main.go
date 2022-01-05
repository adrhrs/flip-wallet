package main

import (
	"log"
	"net/http"
	"sync"

	balanceDelivery "github.com/flip-clean/balance/delivery"
	balanceMemoryRepo "github.com/flip-clean/balance/repository/memory"
	balanceUsecase "github.com/flip-clean/balance/usecase"
	"github.com/flip-clean/models"
	userDelivery "github.com/flip-clean/user/delivery"
	userMemoryRepo "github.com/flip-clean/user/repository/memory"
	userUsecase "github.com/flip-clean/user/usecase"
)

type Dependency struct {
	UserData             map[string]string
	TopTransactionByUser map[string][]models.UserTransaction
	UserDebitValue       map[string]int
	TopUserByTransaction []models.UserDebitTransaction

	BalanceDataSynced sync.Map
}

var (
	userUsecaseSingleton userUsecase.UserUsecase
)

func initInMemoryData() (
	userData map[string]string,
	topTransactionByUser map[string][]models.UserTransaction,
	userDebitValue map[string]int,
	topUserByTransaction []models.UserDebitTransaction,
	balanceDataSynced sync.Map,
) {
	userData = make(map[string]string)
	topTransactionByUser = make(map[string][]models.UserTransaction)
	userDebitValue = make(map[string]int)

	topTransactionByUser["top-transaction-adr"] = []models.UserTransaction{
		{
			Amount:   1000,
			Username: "adr",
		},
		{
			Amount:   900,
			Username: "adr",
		},
		{
			Amount:   -880,
			Username: "adr",
		},
		{
			Amount:   700,
			Username: "adr",
		},
		{
			Amount:   -540,
			Username: "adr",
		},
	}

	topUserByTransaction = []models.UserDebitTransaction{
		{
			Username:        "ppp",
			TransactedValue: 10000,
		},
		{
			Username:        "ooo",
			TransactedValue: 10000,
		},
		{
			Username:        "qqq",
			TransactedValue: 100,
		},

		{
			Username:        "www",
			TransactedValue: 150,
		},
	}

	return
}

func initDependency() Dependency {
	userData, topTransactionByUser, userDebitValue, topUserByTransaction, balanceDataSynced := initInMemoryData()
	return Dependency{
		UserData:             userData,
		TopTransactionByUser: topTransactionByUser,
		UserDebitValue:       userDebitValue,
		TopUserByTransaction: topUserByTransaction,
		BalanceDataSynced:    balanceDataSynced,
	}
}

func initUserUsecase(dep Dependency) {
	userMemoryRepo := userMemoryRepo.NewUserMemoryRepo(dep.UserData)
	userUsecaseSingleton = userUsecase.NewUserUsecase(userMemoryRepo)
	userDelivery.NewUserHandler(userUsecaseSingleton)
}

func initBalanceUsecase(dep Dependency) {
	balanceMemoryRepo := balanceMemoryRepo.NewBalanceMemoryRepo(dep.TopTransactionByUser, dep.UserDebitValue, dep.TopUserByTransaction, dep.BalanceDataSynced)
	balanceUsecase := balanceUsecase.NewBalanceUsecase(balanceMemoryRepo, userUsecaseSingleton)
	balanceDelivery.NewBalanceHandler(balanceUsecase)
}

func main() {

	dep := initDependency()
	initUserUsecase(dep)
	initBalanceUsecase(dep)

	log.Println("server up and running")
	if err := http.ListenAndServe(":4444", nil); err != nil {
		log.Fatal(err)
	}

}
