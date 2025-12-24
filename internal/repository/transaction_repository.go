package repository

import (
	"FinalTaskEvermoss/internal/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx entity.Transaction) error
	GetByID(id uint) (entity.Transaction, error)
	GetByUserID(userID uint) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(tx entity.Transaction) error {
	return r.db.Transaction(func(dbTx *gorm.DB) error {
		// 1. Simpan Transaksi
		if err := dbTx.Create(&tx).Error; err != nil {
			return err
		}
		// 2. Kurangi Stok Produk (Poin Plus!)
		return dbTx.Model(&entity.Product{}).Where("id = ?", tx.ProductID).
			Update("stok", gorm.Expr("stok - ?", tx.Quantity)).Error
	})
}

func (r *transactionRepository) GetByID(id uint) (entity.Transaction, error) {
	var tx entity.Transaction
	err := r.db.Preload("Product").Preload("Address").First(&tx, id).Error
	return tx, err
}

func (r *transactionRepository) GetByUserID(userID uint) ([]entity.Transaction, error) {
	var txs []entity.Transaction
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&txs).Error
	return txs, err
}
