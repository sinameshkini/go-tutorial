package usecase

import (
	"errors"
	"go_tutorial/internal/domain"
	"go_tutorial/internal/repository"
)

type AuthUsecase struct {
	repo *repository.UserRepository
}

func New(repo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (u *AuthUsecase) SignUp(email string, password string) (err error) {
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return err
	}
	if usr != nil {
		return errors.New("Email already exists")
	}

	err = u.repo.AddUser(&domain.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}
	return
}

// TODO impl other methods here ...
