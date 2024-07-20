package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My first goroutine")
	// go func() {
	// 	fmt.Println("My other goroutine")
	// }()

	go firstPrint()
	go secondPrint()

	time.Sleep(time.Millisecond * 5) // akan melihat apakah ada go routine lain yang bisa dieksekusi
	// go others()
}

func others() {
	fmt.Println("My other goroutine")
}

func firstPrint() {
	fmt.Println("My toher go routines!")
}

func secondPrint() {
	fmt.Println("My toher go routines v2.0!")
}