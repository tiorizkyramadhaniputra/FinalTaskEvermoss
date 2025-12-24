package handler

import (
	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/usecase"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	usc usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usc: u}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var prod entity.Product
	json.NewDecoder(r.Body).Decode(&prod)

	if err := h.usc.CreateProduct(userID, prod); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk berhasil ditambah"})
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, _ := h.usc.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}
