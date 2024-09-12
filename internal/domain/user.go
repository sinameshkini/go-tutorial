package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
}

type Email struct {
	Email string
}
