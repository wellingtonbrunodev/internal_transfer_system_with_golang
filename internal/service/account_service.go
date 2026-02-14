package service

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type AccountService struct {
	repo AccountRepositoryInterface
}

func NewAccountService(repo AccountRepositoryInterface) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(ctx context.Context, id int64, initialBalance string) error {
	if id <= 0 {
		return domain.ErrInvalidAccountId
	}

	amount, err := decimal.NewFromString(initialBalance)
	if err != nil {
		return domain.ErrInvalidAmountFormat
	}

	if amount.IsNegative() {
		return domain.ErrInvalidAmountValue
	}

	return s.repo.Create(ctx, id, amount)
}

func (s *AccountService) GetAccount(ctx context.Context, id int64) (int64, decimal.Decimal, error) {
	if id <= 0 {
		return 0, decimal.Zero, domain.ErrInvalidAccountId
	}

	return s.repo.GetByID(ctx, id)
}

