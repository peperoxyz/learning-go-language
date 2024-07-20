package main

import (
	"fmt"
	"sync"
)

func main() {
	// define waitgroup
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		// use wg.Add(2): cause we have 2 goroutines
		wg.Add(2)

		// first goroutine
		go func(index int) {
			defer wg.Done() // to decrement wg counter by 1
			data := []string{"coba1", "coba2", "coba3"}
			fmt.Println(data, index) // print data with index

		}(i)

		// 2nd goroutine
		go func(index int) {
			defer wg.Done()
			data := []string{"bisa1", "bisa2", "bisa3"}
			fmt.Println(data, index)

		}(i)
	}

	// waitgroup: will blocks untill the waitgroup counter is 0 and then stop
	wg.Wait()
}