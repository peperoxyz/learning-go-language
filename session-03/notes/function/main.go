package main

import (
	"fmt"
	"math"
)

func main() {
	// Function
	sum(100,1)
	greeting("Dea", "Haloooo")
	
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println(area, circumference)
}

func sum(a, b int) {
	var sum int
	sum = a + b
	fmt.Println(sum)
}

func greeting(name string, greeting string) {
	fmt.Println(greeting, name)
}

// #1 Fungsi multiple return
func calculate(d float64) (float64, float64) {
	// calculate area
	var area = math.Pi * math.Pow(d / 2, 2)

	// calculate circumference
	var circumference = math.Pi * d

	// return 2 values
	return area, circumference
}

// #2



