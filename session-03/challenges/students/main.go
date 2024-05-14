package main

import (
	"fmt"
	"os"
)

type Student struct {
	id		string
	name 	string
	address string
	reason	string
}

func main() {
	// check if the program is run with CLI arguments
	if len(os.Args) > 1 {
		detailStudentByIdOrName(os.Args[1])
		return
	}
	
	listMenus()

	var menu int
	fmt.Print("\nSelect Menu: ")
	fmt.Scan(&menu);

	switch menu {
	case 1:
		listStudents()
	case 2:
		var idOrName string
		fmt.Print("Enter student ID or name: ")
		fmt.Scan(&idOrName)
		detailStudentByIdOrName(idOrName)
	case 3:
		if len(os.Args) > 1 {
			detailStudentByIdOrName(os.Args[1])
		} else {
			fmt.Println("Please provide student ID or name.\n\n[Ex: go run main.go Dea]")
		}
	default:
		fmt.Println("Unrecognized option.")
	}
}

func listMenus() {
	fmt.Println("\n=========== GFW DATABASE ===========\n1. List of Students \n2. Detail Student By Name \n3. Detail Student By CLI Arguments \n====================================")
}

func listStudents() {
	fmt.Println("\n===== List Students ===== \n1. Dea \n2. Ananda \n3. Gunawan")
	var cont string
		fmt.Print("\nContinue Action (y/n): ")
		fmt.Scan(&cont);
		if cont == "y" {
			main()
		} else if cont == "n" {
			fmt.Println("\nEnd of Service. Thanks!")
		} else {
			fmt.Println("\nUnrecognized option.")
		}
}

func detailStudentByIdOrName(idOrName string) {
	fmt.Print("\n")

	var gfw_students = []Student {
		{id: "1", name: "Dea", address: "Batam, Kepulauan Riau", reason: "Ter-doktrin techbros di X"},
		{id: "2", name: "Ananda", address: "Serang, Banten", reason: "Katanya bagus"},
		{id: "3", name: "Gunawan", address: "Serang, Banten", reason: "Bosen nge-fix xampp yang error mulu"},
	}

	// check if name or ID exists in the list
	var found bool
	for _, student := range gfw_students {
		if student.id == idOrName || student.name == idOrName  {
			fmt.Printf("ID : %v\nName : %v\nAddress : %v\nReason : %v\n", student.id, student.name, student.address, student.reason)
			found = true
		}
	}
	if !found {
		fmt.Println("[404 NOT FOUND] Name or ID does not exist.")
	}
	
}