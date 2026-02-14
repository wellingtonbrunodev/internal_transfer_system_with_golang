package service

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type TransactionService struct {
	repo TransactionRepositoryInterface
}

func NewTransactionService(repo TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Transfer(
	ctx context.Context,
	sourceID int64,
	destinationID int64,
	amount string,
) error {

	if sourceID <= 0 || destinationID <= 0 {
		return domain.ErrInvalidAccountId
	}

	if sourceID == destinationID {
		return domain.ErrSameAccountTransfer
	}

	decimalAmount, err := decimal.NewFromString(amount)
	if err != nil {
		return domain.ErrInvalidAmountFormat
	}

	if decimalAmount.IsNegative() {
		return domain.ErrInvalidAmountValue
	}

	return s.repo.Transfer(ctx, sourceID, destinationID, decimalAmount)
}
