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
	return repo.db.Create(user).Error
}

func (repo *UserRepository) FindUser(email string) (*domain.User, error) {
	var userModel domain.User
	result := repo.db.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		//if the error shows that there is not one in db, than that shoud not be treated as a error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &userModel, nil
}

func (repo *UserRepository) ResetPassword(email string, newPassword string) error {
	var user domain.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}

	user.Password = newPassword
	return repo.db.Save(&user).Error
}
