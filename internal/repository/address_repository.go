package repository

import (
	"FinalTaskEvermoss/internal/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address entity.Address) error
	GetByUserID(userID uint) ([]entity.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address entity.Address) error {
	return r.db.Create(&address).Error
}

func (r *addressRepository) GetByUserID(userID uint) ([]entity.Address, error) {
	var addresses []entity.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
