package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
)

// Book struct with exported fields
type Book struct {
	ID     int    
	Title  string 
	Stock  int    
	Author string 

}

var books = []Book{
	{ID: 1, Title: "Winter in Tokyo", Stock: 20, Author: "Ilana Tan"},
	{ID: 2, Title: "Laskar Pelangi", Stock: 18, Author: "Andrea Hirata"},
}

var PORT = ":2020"

func main() {
	// Register the handlers before starting the server
	http.HandleFunc("/book", createBook)
	http.HandleFunc("/books", getBooks)

	// Start the server
	http.ListenAndServe(PORT, nil)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // bentuk data response yang ingin kita kirim ke client

	// check method
	if r.Method == "POST" {
		title := r.FormValue("title")
		stock := r.FormValue("stock")
		author := r.FormValue("author")

		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
			return
		}

		newBook := Book{
			ID: len(books) + 1,
			Title: title,
			Stock: convertStock,
			Author: author,
		}

		books = append(books, newBook)
		// Return the newly created book as JSON
		json.NewEncoder(w).Encode(newBook)
		return

	}
	
	http.Error(w, "Invalid method", http.StatusBadRequest)
	
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	// check method
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, books)
	}

	// http.Error(w, "Invalid method", http.StatusBadRequest)
}
