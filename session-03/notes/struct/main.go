package main

import "fmt"

// Struct: tipe data yang berupa kumpulan data dari berbagai tipe data
// Struct Employee: Tipe data Employee
	type Employee struct {
		name     string
		age      int
		position string
	}

func main() {
	
	// Giving value to a struct: melalui variabel yang terkait dengan structnya
	var employee Employee
	employee.name = "Dea"
	employee.age = 22
	employee.position = "Junior IT Support"
	
	fmt.Println(employee)

	// Initializing struct sekalian giving valuenya
	var employee1 = Employee{}
	employee1.name = "Adil"
	employee1.age = 19
	employee1.position = "Student"

	fmt.Println(employee1)

	var employee2 = Employee{name: "Gunawan", age: 55, position: "CEO"}
	fmt.Println(employee2)

	// Menggunakan %v untuk memformat struct menjadi string
	fmt.Printf("Employee 1: %v \n", employee1)
	fmt.Printf("Employee 2: %v \n", employee2)

	// Pointer to Struct
	var employee3 *Employee = &employee2 // variabel employee3 dengan tipe data pointer to a struct yang menyimpan alamat memori dari variabel employee2

	fmt.Println(employee3)

	employee3.name = "Dwi" // var employee2 juga terganti, karena berbagi 1 memory address yang sama

	fmt.Printf("Employee 2: %v \n", employee2)
	fmt.Printf("Employee 3: %v \n", employee3)

	
	
}