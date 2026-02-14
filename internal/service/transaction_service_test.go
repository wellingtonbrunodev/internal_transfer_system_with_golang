package service

import (
	"context"
	// "errors"
	"testing"

	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type MockTransactionRepository struct {
	TransferFunc func(ctx context.Context, sourceID, destinationID int64, amount string) error
}

func (m *MockTransactionRepository) Transfer(
	ctx context.Context,
	sourceID int64,
	destinationID int64,
	amount string,
) error {
	return m.TransferFunc(ctx, sourceID, destinationID, amount)
}


func TestTransfer_InvalidAccountID(t *testing.T) {
	mockRepo := &MockTransactionRepository{}
	service := NewTransactionService(mockRepo)

	err := service.Transfer(context.Background(), 0, 2, "100")

	if err != domain.ErrInvalidAccountId {
		t.Errorf("expected ErrInvalidAccountId, got %v", err)
	}
}


func TestTransfer_SameAccount(t *testing.T) {
	mockRepo := &MockTransactionRepository{}
	service := NewTransactionService(mockRepo)

	err := service.Transfer(context.Background(), 1, 1, "100")

	if err != domain.ErrSameAccountTransfer {
		t.Errorf("expected ErrSameAccountTransfer, got %v", err)
	}
}


func TestTransfer_InvalidAmountFormat(t *testing.T) {
	mockRepo := &MockTransactionRepository{}
	service := NewTransactionService(mockRepo)

	err := service.Transfer(context.Background(), 1, 2, "abc")

	if err != domain.ErrInvalidAmountFormat {
		t.Errorf("expected ErrInvalidAmountFormat, got %v", err)
	}
}


func TestTransfer_InvalidAmountValue(t *testing.T) {
	mockRepo := &MockTransactionRepository{}
	service := NewTransactionService(mockRepo)

	err := service.Transfer(context.Background(), 1, 2, "-10")

	if err != domain.ErrInvalidAmountValue {
		t.Errorf("expected ErrInvalidAmountValue, got %v", err)
	}
}

func TestTransfer_Success(t *testing.T) {
	mockRepo := &MockTransactionRepository{
		TransferFunc: func(ctx context.Context, sourceID, destinationID int64, amount string) error {
			return nil
		},
	}

	service := NewTransactionService(mockRepo)

	err := service.Transfer(context.Background(), 1, 2, "100")

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}
