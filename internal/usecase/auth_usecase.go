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
		return errors.New("email already exists")
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

func (u *AuthUsecase) SignIn(email string, password string) (err error) {
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return err
	}
	if usr == nil {
		return errors.New("user not found")
	}

	if usr.Password != password {
		return errors.New("password doesn't match")
	}
	return nil
}

func (u *AuthUsecase) ResetPassowrd(email string) (err error) {
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return err
	}
	if usr == nil {
		return errors.New("user not found")
	}

	newPass := "12ahfhu259"
	if err := u.repo.ResetPassword(usr.Email, newPass); err != nil {
		return err
	}

	return
}
