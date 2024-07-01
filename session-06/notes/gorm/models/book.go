package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Stock     int `gorm:"not null"`
	UserID    uint   // otomatis foreignKey karena User(Struct di model lain)+ID(PK)
	CreatedAt time.Time
	UpdatedAt time.Time
}

// hooks dibuat di model. karena akan merefer ke suatu struct
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book before create()")

	if len(b.Title) < 5 {
		err = errors.New("Book title is too short")
	}

	return
}