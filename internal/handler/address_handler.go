package handler

import (
	"encoding/json"
	"net/http"

	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/usecase"
)

type AddressHandler struct {
	usecase usecase.AddressUsecase
}

func NewAddressHandler(u usecase.AddressUsecase) *AddressHandler {
	return &AddressHandler{usecase: u}
}

func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var address entity.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(userID, address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Alamat berhasil ditambahkan",
	})
}

func (h *AddressHandler) GetMyAddresses(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	addresses, err := h.usecase.GetMyAddresses(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(addresses)
}
