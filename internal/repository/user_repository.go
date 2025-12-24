package repository

import (
	"FinalTaskEvermoss/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user entity.User, toko entity.Toko) error
	GetByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user entity.User, toko entity.Toko) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		toko.UserID = user.ID
		return tx.Create(&toko).Error
	})
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
