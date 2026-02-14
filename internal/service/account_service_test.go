package service

import (
	"context"
	"testing"
	"github.com/shopspring/decimal"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type MockAccountRepository struct {
	CreateFunc  func(ctx context.Context, id int64, balance decimal.Decimal) error
	GetByIDFunc func(ctx context.Context, id int64) (int64, decimal.Decimal, error)
}

func (m *MockAccountRepository) Create(ctx context.Context, id int64, balance decimal.Decimal) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, id, balance)
	}
	return nil
}

func (m *MockAccountRepository) GetByID(ctx context.Context, id int64) (int64, decimal.Decimal, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return 0, decimal.Zero, nil
}


func TestCreateAccount_InvalidID(t *testing.T) {
	mockRepo := &MockAccountRepository{}
	service := NewAccountService(mockRepo)

	err := service.CreateAccount(context.Background(), 0, "100")

	if err != domain.ErrInvalidAccountId {
		t.Errorf("expected ErrInvalidAccountId, got %v", err)
	}
}

func TestCreateAccount_InvalidBalanceFormat(t *testing.T) {
	mockRepo := &MockAccountRepository{}
	service := NewAccountService(mockRepo)

	err := service.CreateAccount(context.Background(), 1, "abc")

	if err != domain.ErrInvalidAmountFormat {
		t.Errorf("expected ErrInvalidAmountFormat, got %v", err)
	}
}

func TestCreateAccount_InvalidNegativeBalance(t *testing.T) {
	mockRepo := &MockAccountRepository{}
	service := NewAccountService(mockRepo)

	err := service.CreateAccount(context.Background(), 1, "-50")

	if err != domain.ErrInvalidAmountValue {
		t.Errorf("expected ErrInvalidAmountValue, got %v", err)
	}
}


func TestCreateAccount_Success(t *testing.T) {
	called := false

	mockRepo := &MockAccountRepository{
		CreateFunc: func(ctx context.Context, id int64, balance decimal.Decimal) error {
			called = true

			if id != 1 {
				t.Errorf("expected id 1, got %d", id)
			}

			if !balance.Equal(decimal.NewFromInt(100)){
				t.Errorf("expected balance 100, got %s", balance)
			}

			return nil
		},
	}

	service := NewAccountService(mockRepo)

	err := service.CreateAccount(context.Background(), 1, "100")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !called {
		t.Error("expected repository Create to be called")
	}
}


func TestGetAccount_InvalidID(t *testing.T) {
	mockRepo := &MockAccountRepository{}
	service := NewAccountService(mockRepo)

	_, _, err := service.GetAccount(context.Background(), 0)

	if err != domain.ErrInvalidAccountId {
		t.Errorf("expected ErrInvalidAccountId, got %v", err)
	}
}

func TestGetAccount_Success(t *testing.T) {
	mockRepo := &MockAccountRepository{
		GetByIDFunc: func(ctx context.Context, id int64) (int64, decimal.Decimal, error) {

			if id != 1 {
				t.Errorf("expected id 1, got %d", id)
			}

			return 1, decimal.NewFromInt(500), nil
		},
	}

	service := NewAccountService(mockRepo)

	id, balance, err := service.GetAccount(context.Background(), 1)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if id != 1 {
		t.Errorf("expected id 1, got %d", id)
	}

	if !balance.Equal(decimal.NewFromInt(500)) {
		t.Errorf("expected balance 500, got %s", balance)
	}
}

func TestGetAccount_RepoError(t *testing.T) {
	mockRepo := &MockAccountRepository{
		GetByIDFunc: func(ctx context.Context, id int64) (int64, decimal.Decimal, error) {
			return 0, decimal.Zero, domain.ErrAccountNotFound
		},
	}

	service := NewAccountService(mockRepo)

	_, _, err := service.GetAccount(context.Background(), 1)

	if err != domain.ErrAccountNotFound {
		t.Errorf("expected ErrAccountNotFound, got %v", err)
	}
}
