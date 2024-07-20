package main

import (
	"fmt"
	"sync"
	"time"
)

/* waitgroup: struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutine.

karena menggunakan time.sleep untuk menjalankan goroutine tidak terlalu efektif, maka dibutuhkan sesuatu untuk mengoptimalisasi penggunaan goroutine.

*/

func main() {
	// define waitgroup
	var wg sync.WaitGroup

	// use waitgroup: cause we have 2 goroutine
	wg.Add(2)
	
	go printName("Dea", &wg)
	go printName("Dwi", &wg)
	// waitgroup: will blocks untill the waitgroup counter is 0 and then stop
	wg.Wait()

	// time.Sleep(time.Second * 5)
}

func printName(s string, wg *sync.WaitGroup) {
	for i:=0; i<5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	defer wg.Done()
}