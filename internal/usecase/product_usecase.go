package usecase

import (
	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(userID uint, product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo}
}

func (u *productUsecase) CreateProduct(userID uint, product entity.Product) error {
	// Cari toko milik user yang login
	toko, err := u.repo.GetTokoByUserID(userID)
	if err != nil {
		return err
	}
	product.TokoID = toko.ID // Set TokoID otomatis
	return u.repo.Create(product)
}

func (u *productUsecase) GetAllProducts() ([]entity.Product, error) {
	return u.repo.GetAll()
}
