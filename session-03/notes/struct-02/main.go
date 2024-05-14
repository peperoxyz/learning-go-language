package main

import (
	"fmt"
)

// Embeded Struct: Struct yang mengandung tipe data struct lainnya di dalamnya
type Person struct {
	name 	string
	age 	int
}

type Member struct {
	position 	string
	person 		Person
}


func main() {
	var member1 = Member{}
	member1.position = "Rapper"
	member1.person.name = "Giselle"
	member1.person.age = 20

	// fmt.Printf("%v", member1)

	// Anonymous Struct: Struct yang tidak dideclare di awal, deklarasi harus diikuti dengan inisialisasi
	// Anonimous struct tanpa pengisian field secara langsung
	var member2 = struct {
		person Person
		position string
	}{}
	member2.person = Person{name: "Ningning", age: 21}
	member2.position = "Lead Vocal"
	
	fmt.Printf("Member 2: %v", member2)

	// Anonymous struct dengan langsung pengisian field
	var member3 = struct {
		person Person
		position string
	}{
		person: Person{name:"Karina", age: 21},
		position: "Visual, Leader",
	}

	fmt.Printf("\nMember 3: %v\n", member3)
	

	// Slice to Struct: Kumpulan data dengan tipe datanya struct tersebut
	var people = []Person {
		{name: "Dea", age: 22},
		{name: "Nilam", age: 24},
		{name: "Kelly", age: 19},
		{name: "Maria", age: 25},
	}


	for _, value:= range people {
		fmt.Printf("\n%v", value)
	}

	// Slice of anonymous struct
	var member = []struct {
		person Person
		position string
	} {
		{person: Person{name: "Jisoo", age: 27}, position: "Visual, vocal"},
		{person: Person{name: "Jennie", age: 27}, position: "Rapper, vocal"},
		{person: Person{name: "Rose", age: 26}, position: "Main vocal"},
		{person: Person{name: "Lisa", age: 26}, position: "Rapper, main dancer"},
	}
	fmt.Print("\n")
	for _, value := range member {
		fmt.Printf("\n%v", value)
	}


}