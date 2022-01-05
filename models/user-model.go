package models

import "errors"

type (
	UserRegistrationPayload struct {
		Username string `json:"username"`
		Token    string
	}

	UserRegistrationResponse struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	}
)

var (
	ErrUserExist                     = errors.New("user already exist")
	ErrUserNotFound                  = errors.New("destination user not found")
	ErrInsufficientBalance           = errors.New("insufficient balance")
	ErrDecSenderBalance              = errors.New("err inc sender balance")
	ErrIncTargetBalance              = errors.New("err dec sender balance")
	ErrSenderSetTopTransactionByUser = errors.New("err set top transaction data by user (sender)")
	ErrTargetSetTopTransactionByUser = errors.New("err set top transaction data by user (target)")
	ErrSetTotalDebitSender           = errors.New("err set total debit sender")
	ErrSetTopUserByTransaction       = errors.New("err set top user by transaction")
)
