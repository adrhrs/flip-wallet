package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/flip-clean/balance/repository"
	"github.com/flip-clean/models"
	"github.com/flip-clean/user/usecase"
)

func Test_balanceUsecase_IncrementUserBalance(t *testing.T) {
	type args struct {
		ctx   context.Context
		param models.IncrementUserBalancePayload
	}
	tests := []struct {
		name    string
		b       *balanceUsecase
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(context.Background(), "username", "adr"),
				param: models.IncrementUserBalancePayload{
					Amount: 1000,
				},
			},
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					IncBalanceFunc: func(username string, increment int) error {
						return nil
					},
				},
			},
		},
		{
			name: "failed",
			args: args{
				ctx: context.WithValue(context.Background(), "username", "adr"),
				param: models.IncrementUserBalancePayload{
					Amount: 1000,
				},
			},
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					IncBalanceFunc: func(username string, increment int) error {
						return errors.New("err")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.IncrementUserBalance(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("balanceUsecase.IncrementUserBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_balanceUsecase_TransferBalance(t *testing.T) {
	type args struct {
		ctx   context.Context
		param models.TransferBalancePayload
	}
	tests := []struct {
		name    string
		b       *balanceUsecase
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(context.Background(), "username", "adr"),
				param: models.TransferBalancePayload{
					ToUsername: "afn",
					Amount:     100,
				},
			},
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					GetBalanceFunc: func(username string) (int, bool) {
						return 1000, true
					},
					IncBalanceFunc: func(username string, increment int) error {
						return nil
					},
					GetTopTransactionByUserFunc: func(username string) []models.UserTransaction {
						return []models.UserTransaction{
							{
								Username: "adr",
								Amount:   200,
							},
						}
					},
					SetTopTransactionByUserFunc: func(username string, value []models.UserTransaction) error {
						return nil
					},
					GetUserDebitValueFunc: func(username string) int {
						return 200
					},
					SetUserDebitValueFunc: func(username string, value int) error {
						return nil
					},
					GetTopUserByTransactionFunc: func() []models.UserDebitTransaction {
						return []models.UserDebitTransaction{
							{
								Username:        "qq",
								TransactedValue: 300,
							},
						}
					},
					SetTopUserByTransactionFunc: func(value []models.UserDebitTransaction) error {
						return nil
					},
				},
				UserUsecase: &usecase.UserUsecaseMock{
					IsUserExistFunc: func(username string) bool {
						return true
					},
				},
			},
		},
		{
			name: "fail err code 6",
			args: args{
				ctx: context.WithValue(context.Background(), "username", "adr"),
				param: models.TransferBalancePayload{
					ToUsername: "afn",
					Amount:     100,
				},
			},
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					GetBalanceFunc: func(username string) (int, bool) {
						return 1000, true
					},
					IncBalanceFunc: func(username string, increment int) error {
						return nil
					},
					GetTopTransactionByUserFunc: func(username string) []models.UserTransaction {
						return []models.UserTransaction{
							{
								Username: "adr",
								Amount:   200,
							},
						}
					},
					SetTopTransactionByUserFunc: func(username string, value []models.UserTransaction) error {
						return nil
					},
					GetUserDebitValueFunc: func(username string) int {
						return 200
					},
					SetUserDebitValueFunc: func(username string, value int) error {
						return nil
					},
					GetTopUserByTransactionFunc: func() []models.UserDebitTransaction {
						return []models.UserDebitTransaction{
							{
								Username:        "qq",
								TransactedValue: 300,
							},
						}
					},
					SetTopUserByTransactionFunc: func(value []models.UserDebitTransaction) error {
						return errors.New("err")
					},
				},
				UserUsecase: &usecase.UserUsecaseMock{
					IsUserExistFunc: func(username string) bool {
						return true
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.TransferBalance(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("balanceUsecase.TransferBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_balanceUsecase_GetTopTransactionByUser(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		b    *balanceUsecase
		args args
		want []models.UserTransaction
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(context.Background(), "username", "adr"),
			},
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					GetTopTransactionByUserFunc: func(username string) []models.UserTransaction {
						return []models.UserTransaction{}
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetTopTransactionByUser(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceUsecase.GetTopTransactionByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_balanceUsecase_GetTopUserByTransaction(t *testing.T) {
	tests := []struct {
		name string
		b    *balanceUsecase
		want []models.UserDebitTransaction
	}{
		{
			name: "success",
			b: &balanceUsecase{
				BalanceRepo: &repository.BalanceRepoMock{
					GetTopUserByTransactionFunc: func() []models.UserDebitTransaction {
						return []models.UserDebitTransaction{}
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetTopUserByTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceUsecase.GetTopUserByTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
