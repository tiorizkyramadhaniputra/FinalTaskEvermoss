package repository

import (
	"FinalTaskEvermoss/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product entity.Product) error
	GetTokoByUserID(userID uint) (entity.Toko, error)
	GetAll() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product entity.Product) error {
	return r.db.Create(&product).Error
}

func (r *productRepository) GetTokoByUserID(userID uint) (entity.Toko, error) {
	var toko entity.Toko
	err := r.db.Where("user_id = ?", userID).First(&toko).Error
	return toko, err
}

func (r *productRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}
