package service

import (
	"context"
	"strconv"
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

	parsedAmount, err := strconv.ParseFloat(initialBalance, 64)

	if err != nil || parsedAmount < 0 {
		return domain.ErrInvalidAmountFormat
	}

	return s.repo.Create(ctx, id, initialBalance)
}

func (s *AccountService) GetAccount(ctx context.Context, id int64) (int64, string, error) {
	if id <= 0 {
		return 0, "", domain.ErrInvalidAccountId
	}

	return s.repo.GetByID(ctx, id)
}

