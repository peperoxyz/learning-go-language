package main

import "fmt"

func main() {
	// FizzBuzz App for Mini Challenge - 1
	fmt.Println("\n=== Welcome to FizzBuzz App ===")
	
	// get the number from user input
	var number int
	fmt.Print("Input a number: ")
	fmt.Scan(&number);

	// looping
	for i := 1; i <= number; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i % 3 == 0 {
				fmt.Println("Fizz")
			} else if i % 5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
	}
}