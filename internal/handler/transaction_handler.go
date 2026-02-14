package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/service"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

type transferRequest struct {
	SourceAccountID      int64  `json:"source_account_id"`
	DestinationAccountID int64  `json:"destination_account_id"`
	Amount               string `json:"amount"`
}

func (h *TransactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req transferRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.Transfer(
		r.Context(),
		req.SourceAccountID,
		req.DestinationAccountID,
		req.Amount,
	)

	if err != nil {

		status, ok := domain.APIStatusMap[err]

		if ok {
			http.Error(w, err.Error(), status)
			return
		}

		http.Error(w, "internal error", http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}
