package main

import "fmt"

func main() {
	// define two channels
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	for i := 1; i < 5; i++ {
		// first goroutine
		go func() {
			data := []string{"coba1", "coba2", "coba3"}
			ch1 <- data // send data to ch1
		}()

		// second goroutine
		go func(index int) {
			result1 := <-ch1 // receive data from ch1
			fmt.Println(result1, i) // print received data with index
			data := []string{"bisa1", "bisa2", "bisa3"}
			ch2 <- data // send data to ch2
		}(i)

		// main goroutine receives data from ch2
		result2 := <-ch2
		fmt.Println(result2, i) // print received data with index
	}
}
