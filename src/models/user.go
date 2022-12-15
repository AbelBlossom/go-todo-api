package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Todos    []Todo `json:"todos"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	if hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10); err == nil {
		user.Password = string(hash)
	} else {
		return err
	}

	return
}

func (user *User) ComparePassword(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err == nil {
		return true
	}
	return false
}
