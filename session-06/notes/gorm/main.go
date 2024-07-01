package main

import (
	"errors"
	"fmt"
	"learning-go-language/session-06/notes/gorm/database"
	"learning-go-language/session-06/notes/gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	// CreateUser("deaanandagunawan@gmail.com")
	// GetUserById(12)
	// UpdateUserById(1, "dea@gmail.co.id")
	CreateBook(1, "Laskar Pelangi", "Andrea Hirata", 40 )
}

func CreateUser(email string) {
	db := database.GetDB() // baru bisa pakai seluruh function di gorm
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	// create record with gorm
	// harus define isi dari struct user itu apa: dalam case ini, model user
	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("Error creating user data.")
		return
	}
	fmt.Println("New user data:", user)
}

func GetUserById(id uint) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	// harus define isi dari struct user itu apa: dalam case ini, model user
	user := models.User{}

	// proses find nya
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("User data not found")
			return
		}
		print("Error finding user: ", err)
	}
	fmt.Printf("User data: %v \n", user)
}

func UpdateUserById(id uint, email string) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func CreateBook(userId uint, title string, author string, stock int) {
	// hooks ditulis di model

	// define db
	db := database.GetDB()

	// define struct
	Book := models.Book{
		UserID: userId,
		Title: title,
		Author: author,
		Stock: stock,
	}
	
	err := db.Create(&Book).Error
	if err != nil {
		fmt.Println("Error creating book data.", err)
		return
	}

	fmt.Println("New book data successfully created: ", Book)
}