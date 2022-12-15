package models

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string
	UserID    uint
	Completed bool `gorm:"default:false"`
}
