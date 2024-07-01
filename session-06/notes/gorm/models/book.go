package models

import "time"

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Stock     string `gorm:"not null"`
	UserID    uint   // otomatis foreignKey karena User(Struct di model lain)+ID(PK)
	CreatedAt time.Time
	UpdatedAt time.Time
}