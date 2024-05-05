package main

import "fmt"

func main() {
	// slice: an array-like data type that is not fixed-length

	// ways to do declare a slice
	// #1 declare and initialize the value
	var cars = []string{"Volvo, Mazda, Toyota"}
	fmt.Println(cars)

	// #2 using make function
	var animals = make([]string,3)
	animals[0] = "Chicken"
	animals[1] = "Fish"
	animals[2] = "Butterfly"

	//  manipulating slice with append function
	animals = append(animals, "Bear", "Panda", "Snake")
	fmt.Println(animals)
	
	// manipulating slice with append func with ellipsis
	fruits1 := []string{"Apple", "Mango", "Peach"}
	fruits2 := []string{"Orange", "Grape", "Strawberry"}

	// appending all element in fruits2 to fruits1
	fruits1 = append(fruits1, fruits2...)
	fmt.Println(fruits1)

	// ==================================================

	// slicing the slice
	menus := []string{"Coffees", "Non-Coffee", "Sandwich", "Pastry", "Alcohol", "Beer"}

	// [x:y] = [start:stop] = [index:urutan-akhir], where index starts from 0, urutan starts from 1
	// i want to get from ["Mon-Coffee" until "Pastry"]
	var menus2 = menus[1:4]
	fmt.Println(menus, menus2)

	// combining slice and append
	colors := []string{"Blue", "Lavender", "Maroon", "Sage"}
	// [:3] means it will get every elements before index 3, and add a new value after that. the output will be : Blue, Lavender, Maroon, + Pink
	colors = append(colors[:3], "pink")
	fmt.Println(colors)
}