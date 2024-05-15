package main

import (
	"fmt"
	"net/http"
)

// Nett/Http Package: package yang menyediakan keperluan dalam membuat app berbasis web: routing, templating, dan lain-lain
// Package templating: package yang menyediakan fitur untuk membuat template HTML

var PORT = ":2020"
	
func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world"
	fmt.Fprintf(w, msg)
}