package domain

import (
	"errors"
	"net/http"
)
var (

	// ErrInsufficientBalance is returned when a transaction is tried with higher amount than account can process.
	ErrInsufficientBalance = errors.New("insufficient balance")
	
	// ErrAccountNotFound is returned when a transaction contains one or more account not present in DB.
	ErrAccountNotFound     = errors.New("one or both accounts not found")

	// ErrSameAccountTransfer is returned when a transaction tries to transfer money to same account.
	ErrSameAccountTransfer = errors.New("cannot transfer to same account")

	// ErrInvalidAccountId is returned when trying to create account with ID < 0.
	ErrInvalidAccountId = errors.New("invalid account id")

	// ErrInvalidAccountId is returned when the amount sent to API is not numeric
	ErrInvalidAmountFormat = errors.New("invalid amount format")

	// ErrInvalidAccountId is returned when the amount sent to API for transaction is less <= 0
	ErrInvalidAmountValue = errors.New("invalid amount value")

	// APIStatusMap maps the API errors to the proper http status
	APIStatusMap = map[error]int{
		ErrInsufficientBalance: http.StatusBadRequest,   
		ErrAccountNotFound:   http.StatusNotFound,   
		ErrSameAccountTransfer:   http.StatusBadRequest,
		ErrInvalidAccountId: http.StatusBadRequest,
		ErrInvalidAmountFormat: http.StatusBadRequest,
		ErrInvalidAmountValue: http.StatusBadRequest,
	}
)

