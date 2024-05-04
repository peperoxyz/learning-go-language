package main

import (
	"fmt"
)

func main() {
	// Letter Counter App for Mini Challenge - 2
	fmt.Println("\n=== Welcome to LetterCounter App ===")
	
	var word string
	fmt.Print("Input the word: ")
	fmt.Scan(&word);

	// create the map
	var letterCounts = make(map[string]int)
	for i := 0; i < len(word); i++ {
		letters := string(word[i])
		println(letters)
		letterCounts[letters]++
	}
	fmt.Println(letterCounts)
}