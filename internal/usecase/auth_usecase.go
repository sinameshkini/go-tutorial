package usecase

import (
	"errors"
	"go_tutorial/internal/domain"
	"go_tutorial/internal/repository"
	"math/rand"
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

func (u *AuthUsecase) ResetPassowrd(email string) (newpass string, err error) {
	var newPass string
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return newPass, err
	}
	if usr == nil {
		return newPass, errors.New("user not found")
	}

	newPass = generateRandomPassword(12)
	if err := u.repo.ResetPassword(usr.Email, newPass); err != nil {
		return newPass, err
	}

	return newPass, err
}

func generateRandomPassword(lengh int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var passowrd []byte

	for i := 0; i <= lengh; i++ {
		passowrd = append(passowrd, charset[rand.Intn(len(charset))])
	}

	return string(passowrd)

}
