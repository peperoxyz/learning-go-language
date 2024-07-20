package main

import (
	"fmt"
	"math"
)

// interface: tipe data di golang yang merupakan kumpulan dari definisi 1 atau lebih method
/* type namaInterface interface {
	method1() returnType
	method2() returnType
}
*/

// define interface
type hitung interface {
	luas() float64
	keliling() float64
} 

// define struct
type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter, jariJari float64
}

// define method
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJariMethod(), 2)
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// function yang hanya bisa diimplement oleh lingkaran saja: mendapatkan nilai jari-jari dari diameter lingkaran
func (l lingkaran) jariJariMethod() float64 {
	return l.diameter/2
}

func main () {
	var lingkaran1 hitung = lingkaran{14, 7}
	var persegi1 hitung = persegi{8}

	fmt.Printf("Type of lingkaran: %T \n", lingkaran1)
	fmt.Printf("Type of persegi: %T \n", persegi1)
	fmt.Println("================================")

	// akses
	lingkaran1.luas()
	lingkaran1.keliling()
	fmt.Println("Luas lingkaran1: ", lingkaran1.luas())
	fmt.Println("Keliling lingkaran1: ", lingkaran1.keliling())
	fmt.Println("================================")
	fmt.Println("Luas persegi1: ", persegi1.luas())
	fmt.Println("Keliling persegi1: ", persegi1.keliling())

	/*
		yang tidak bisa dilakukan:
		lingkaran1.jariJariMethod(); // lingkaran1.jariJariMethod undefined (type hitung has no field or method jariJariMethod)

		maka dibutuhkanlah type assertion: konversi tipe data interface ke tipe data aslinya
	*/

	// yang bisa dilakukan agar lingkaran1 dapat mengakses method jariJariMethod()
	lingkaran1.(lingkaran).jariJariMethod();


}
