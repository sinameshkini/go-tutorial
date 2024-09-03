package repository

import (
	"errors"
	"gu_tutorial/interval/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func StartDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&domain.UserModel{})
	return db, err
}

func (repo *UserRepository) AddUser(user *domain.UserModel) error {
	newUser := domain.UserModel{Email: user.Email, Password: user.Password}
	err := repo.db.Create(&newUser).Error
	return err
}

func (repo *UserRepository) FindUser(email string) (*domain.UserModel, error) {
	var userModel domain.UserModel
	result := repo.db.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &userModel, nil
}
