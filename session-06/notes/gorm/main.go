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
	// CreateUser("dwiyana@gmail.com")
	// GetUserById(12)
	// UpdateUserById(1, "dea@gmail.co.id")
	// CreateBook(2, "Laskar Pelangi", "Andrea Hirata", 40 )
	GetUserWithBooks()
	// DeleteBookById(1)
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

func GetUserWithBooks() {
	db := database.GetDB()
	users := models.User{}

	// preloading (join table book yang terhubung)
	err := db.Preload("Books").Find(&users).Error
	if err != nil {
		fmt.Println("Error getting user data with books.", err.Error())
		return
	}

	fmt.Println("User datas with books: ")
	fmt.Printf("%+v", users)
}

func DeleteBookById(id uint) {
	db := database.GetDB()

	book := models.Book{}

	// check the book is exist or not
	errFind := db.First(&book, id).Error
	if errFind != nil {
		if errors.Is(errFind, gorm.ErrRecordNotFound){
			fmt.Println("Book with id", id, "not found.")
			return
		}
		fmt.Println("Error finding book:", errFind.Error())
		return
	}

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		print("Error deleting book: ", err.Error())
		return
	}

	fmt.Printf("Book with id %d  has been successfully deleted.", id)
}