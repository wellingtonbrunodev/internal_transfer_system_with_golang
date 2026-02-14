package service

import (
	"context"
	"strconv"

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

	parsedAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return domain.ErrInvalidAmountFormat
	}

	if parsedAmount <= 0 {
		return domain.ErrInvalidAmountValue
	}

	return s.repo.Transfer(ctx, sourceID, destinationID, amount)
}
