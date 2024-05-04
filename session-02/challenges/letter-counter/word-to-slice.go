package main

import (
	"fmt"
)

func main2() {
	// Letter Counter App for Mini Challenge - 1
	fmt.Println("\n=== Welcome to LetterCounter App ===")
	
	// get the word from user input
	var word string
	fmt.Print("Input the word: ")
	fmt.Scan(&word);

	var letters = make([]string, len(word))
	for i := 0; i < len(word); i++ {
		letters[i] =  string(word[i])
	}

	fmt.Println(letters)
}