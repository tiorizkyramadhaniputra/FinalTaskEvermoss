package usecase

import (
	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/repository"
	"errors"
)

type TransactionUsecase interface {
	Checkout(userID uint, input entity.Transaction) (entity.Transaction, error)
	GetMyHistory(userID uint) ([]entity.Transaction, error)
}

type transactionUsecase struct {
	txRepo   repository.TransactionRepository
	prodRepo repository.ProductRepository
}

func NewTransactionUsecase(tr repository.TransactionRepository, pr repository.ProductRepository) TransactionUsecase {
	return &transactionUsecase{tr, pr}
}

func (u *transactionUsecase) Checkout(userID uint, input entity.Transaction) (entity.Transaction, error) {
	// 1. Cek Produk & Stok
	var products, _ = u.prodRepo.GetAll() // Simple check for example
	var targetProduct entity.Product
	for _, p := range products {
		if p.ID == input.ProductID {
			targetProduct = p
		}
	}

	if targetProduct.ID == 0 || targetProduct.Stok < input.Quantity {
		return entity.Transaction{}, errors.New("produk tidak ditemukan atau stok kurang")
	}

	// 2. Hitung Total Harga
	input.UserID = userID
	input.TotalPrice = targetProduct.Harga * input.Quantity
	input.Status = "SUCCESS"

	// 3. Simpan
	err := u.txRepo.Create(input)
	return input, err
}

func (u *transactionUsecase) GetMyHistory(userID uint) ([]entity.Transaction, error) {
	return u.txRepo.GetByUserID(userID)
}
