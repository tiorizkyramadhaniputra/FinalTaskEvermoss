package usecase

import (
	"errors"

	"FinalTaskEvermoss/internal/entity"
	"FinalTaskEvermoss/internal/repository"
	"FinalTaskEvermoss/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Registration(user entity.User) error
	Login(email, password string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Registration(user entity.User) error {
	toko := entity.Toko{
		NamaToko: "Toko " + user.Nama,
	}
	return u.repo.Register(user, toko)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("email tidak terdaftar")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("password salah")
	}

	return utils.GenerateToken(user.ID, user.IsAdmin)
}
