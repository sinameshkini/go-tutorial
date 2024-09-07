package repository

import (
	"errors"
	"go_tutorial/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUser(user *domain.User) error {
	newUser := domain.User{Email: user.Email, Password: user.Password}
	err := repo.db.Create(&newUser).Error
	return err
}

func (repo *UserRepository) FindUser(email string) (*domain.User, error) {
	var userModel domain.User
	result := repo.db.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &userModel, nil
}
