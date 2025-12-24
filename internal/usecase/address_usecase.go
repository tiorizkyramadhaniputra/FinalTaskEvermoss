package usecase

import (
	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/repository"
)

type AddressUsecase interface {
	Create(userID uint, address entity.Address) error
	GetMyAddresses(userID uint) ([]entity.Address, error)
}

type addressUsecase struct {
	repo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecase{repo}
}

func (u *addressUsecase) Create(userID uint, address entity.Address) error {
	address.UserID = userID
	return u.repo.Create(address)
}

func (u *addressUsecase) GetMyAddresses(userID uint) ([]entity.Address, error) {
	return u.repo.GetByUserID(userID)
}
