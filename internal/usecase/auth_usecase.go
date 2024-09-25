package usecase

import (
	"errors"
	"go_tutorial/internal/domain"
	"go_tutorial/internal/repository"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
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
	//trying to check if the user already exists
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return err
	}
	if usr != nil {
		return errors.New("email already exists")
	}

	//hashing passowrd for storing
	hashedPassowrd, err := HashPassword(password)
	if err != nil {
		return errors.New("error while hashing passowrd")
	}

	//adding to the database
	err = u.repo.AddUser(&domain.User{
		Email:    email,
		Password: hashedPassowrd,
	})
	if err != nil {
		return err
	}
	return
}

func (u *AuthUsecase) SignIn(email string, password string) (err error) {
	//checking if the user exists
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return err
	}
	if usr == nil {
		return errors.New("user not found")
	}

	//checking if the hashed password and real password match
	if !CheckPasswordHash(password, usr.Password) {
		return errors.New("password doesn't match")
	}
	return nil
}

func (u *AuthUsecase) ResetPassowrd(email string) (newpass string, err error) {

	var newPass string

	//checking if the user exists
	usr, err := u.repo.FindUser(email)
	if err != nil {
		return newPass, err
	}
	if usr == nil {
		return newPass, errors.New("user not found")
	}

	//generating random password and hashing it
	newPass = generateRandomPassword(12)
	newHashedPass, err := HashPassword(newPass)

	if err != nil {
		return newPass, errors.New("error while hashing")
	}

	//trying to change password in db
	if err := u.repo.ResetPassword(usr.Email, newHashedPass); err != nil {
		return newPass, err
	}

	return newPass, err
}

func generateRandomPassword(lengh int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password []byte
	for i := 0; i <= lengh; i++ {
		password = append(password, charset[rand.Intn(len(charset))])
	}
	return string(password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
