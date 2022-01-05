package models

type (
	IncrementUserBalancePayload struct {
		Amount int `json:"amount"`
	}
	IncrementUserBalanceResponse struct {
		Message string `json:"message"`
	}
	GetUserBalanceInfoResponse struct {
		Balance int `json:"balance"`
	}
	TransferBalancePayload struct {
		Amount     int    `json:"amount"`
		ToUsername string `json:"to_username"`
	}
	UserTransaction struct {
		Amount   int    `json:"amount"`
		Username string `json:"username"`
	}

	UserDebitTransaction struct {
		TransactedValue int    `json:"transacted_value"`
		Username        string `json:"username"`
	}
)
