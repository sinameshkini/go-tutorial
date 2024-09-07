package domain

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Email    string
	Password string
}
