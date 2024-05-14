package main

import (
	"fmt"
	"strings"
)

func main() {
	// #1 Pointer: Variabel yang digunakan untuk menyimpan alamat memori variabel lainnya

	/*
		// Pointer dipakai saat:
		   1. Menghindari penggandaan data
		   2. Nilai kosong yang bermakna
		   3. Mengubah nilai variabel di luar fungsi
	*/

	// Deklarasi pointer: var var_name *var-type
	// var firstExample *int
	// var secondExample *string

	var firstNumber int = 4
	var secondNumber *int = &firstNumber // variabel pointer int yang menampung alamat memori dari var firstNumber

	fmt.Println("firstNumber (value): ", firstNumber)
	fmt.Println("firstNumber (memory address): ", &firstNumber)
	fmt.Println("another way of firstNumber's (memory address): ", secondNumber)
	fmt.Println("secondNumber (value) == firstNumber (value): ", *secondNumber)
	fmt.Println("secondNumber (memory address): ", &secondNumber)
	
	// #2 Changing value through pointer
	var firstGroup string = "EXO"
	var secondGroup *string = &firstGroup
	fmt.Println("\n", strings.Repeat("-", 20), "\n")

	fmt.Println("firstGroup (value):", firstGroup )
	fmt.Println("firstGroup (memory address):", &firstGroup)
	fmt.Println("secondGroup (value):", *secondGroup)
	fmt.Println("another way of firstGroup's (memory address):", secondGroup)
	fmt.Println("secondGroup (memory address):", &secondGroup)

	fmt.Println("\n", strings.Repeat("-", 20), "\n")



}