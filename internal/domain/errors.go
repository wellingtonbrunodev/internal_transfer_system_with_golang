package domain

import (
	"errors"
	"net/http"
)
var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrAccountNotFound     = errors.New("one or both accounts not found")
	ErrSameAccountTransfer = errors.New("cannot transfer to same account")
	ErrInvalidAccountId = errors.New("invalid account id")
	ErrInvalidAmountFormat = errors.New("invalid amount format")
	ErrInvalidAmountValue = errors.New("invalid amount value")

	APIStatusMap = map[error]int{
		ErrInsufficientBalance: http.StatusBadRequest,   
		ErrAccountNotFound:   http.StatusNotFound,   
		ErrSameAccountTransfer:   http.StatusBadRequest,
		ErrInvalidAccountId: http.StatusBadRequest,
		ErrInvalidAmountFormat: http.StatusBadRequest,
		ErrInvalidAmountValue: http.StatusBadRequest,
	}
)

