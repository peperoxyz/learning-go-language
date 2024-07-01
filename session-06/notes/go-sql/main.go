package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

const (
	user = "root"
	password = ""
	dbName = "golang_mysql"
)

type Book struct {
	ID     int    
	Title  string 
	Stock  int    
	Author string 
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbName)
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping() // wajib dipanggil
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateBook()
	// GetBooks()
	// UpdateBook()
	// DeleteBook()
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `INSERT INTO books (title, author, stock) VALUES (?, ?, ?)`

	// execute raw query
	result, err := db.Exec(sqlStatement, "Filosofi Teras", "Henry Manampiring", 20)
	if err != nil {
		panic(err)
	}

	// check last insert id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// retreive data
	sqlRetreive := `SELECT * FROM books WHERE id = ?`

	err = db.QueryRow(sqlRetreive, lastInsertId).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data: %+v\n", book)
}

func GetBooks() {
	var results = []Book{}

	sqlStatement := `SELECT * FROM books`

	rows, err := db.Query(sqlStatement)
	if err != nil {
	panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Println("Book data:", results)
}

func DeleteBook() {
	sqlStatement := `DELETE from books WHERE id = ?;`

	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount: ", count)
}

func UpdateBook() {
	sqlStatement := `UPDATE books SET title = ?, author = ?, stock = ? WHERE id =?;`
	res, err := db.Exec(sqlStatement, "Laskar Anak Pulau", "Dwiyana", 30, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}