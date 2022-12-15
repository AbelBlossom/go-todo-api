package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Todos     []Todo    `json:"todos"`
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
