package main

import "fmt"

func main() {
	// defer fmt.Println("1") // akan selalu terakhir diprint karena defer
	// fmt.Println("2")

	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}