package models

import "time"

// models berisi: define data dari table di database

type User struct {
	ID uint				`gorm:"primaryKey"`
	Email string		`gorm:"not null; unique; type:varchar(50)"`
	Books []Book //kebutuhan association
	CreatedAt time.Time
	UpdatedAt time.Time
}

