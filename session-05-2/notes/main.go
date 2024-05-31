package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PORT = ":2020"

type Book struct {
	ID int
	Title string
	Stock int
	Author string
}

var books = []Book {
	{
		ID: 1, 
		Title: "Filosofi Teras", 
		Stock: 29, 
		Author: "Henri Manampiring",
	},
	{
		ID: 2, 
		Title: "Teka Teki Rumah Aneh", 
		Stock: 14, 
		Author: "Uketsu",
	},
}

func main() {
	// keperluan routing
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/book", createBook)
	http.HandleFunc("/books", getBooks)

	http.ListenAndServe(PORT, nil)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	// bentuk data response yang ingin kita kirim ke client
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		// convert data from struct to JSON
		title := r.FormValue("title")
		stock := r.FormValue("stock")
		author := r.FormValue("author")

		// convert the string to int
		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
		}

		// receive data from formData
		newBook := Book {
			ID: len(books) + 1,
			Title: title,
			Stock: convertStock,
			Author: author,
		}

		// append to the slice books
		books = append(books, newBook)

		json.NewEncoder(w).Encode(books)
		return
	}

	// if the method is not get, throw error 
	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	/*
	// return JSON for API response
	// bentuk data response yang ingin kita kirim ke client
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// convert data from struct to JSON
		json.NewEncoder(w).Encode(books)
		return
	}

	// if the method is not get, throw error
	http.Error(w, "Invalid method", http.StatusBadRequest)
	*/

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html") //return tipe data struct template
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		tpl.Execute(w, books)
		return 
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	fmt.Fprintf(w, msg)
}

