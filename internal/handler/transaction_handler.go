package handler

import (
	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/usecase"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	usc usecase.TransactionUsecase
}

func NewTransactionHandler(u usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usc: u}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	var input entity.Transaction
	json.NewDecoder(r.Body).Decode(&input)

	result, err := h.usc.Checkout(userID, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
