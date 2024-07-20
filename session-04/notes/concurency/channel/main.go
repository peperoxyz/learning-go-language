package main

import "fmt"

/*
	channel: mekanisme untuk goroutine saling berkomunikasi dengan goroutine lainnya: proses pengiriman dan pertukaran data antara goroutine satu dengan lainnya
*/

func main() {
	ch := make(chan string) // c adalah variabel yang memiliki tipe data channel of string

	fmt.Println(ch)

}