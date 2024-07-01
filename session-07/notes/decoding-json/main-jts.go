package main

import (
	"encoding/json"
	"fmt"
)

// JSON Decoding to Struct
// Struct student
type People struct {
	FullName string	`json:"full_name"`
	Email string 	`json:"email"`
	Age int			`json:"age"`
}

func mains () {
	var jsonString = `
		{
			"full_name": "Dea Ananda",
			"email": "dea@gmail.com",
			"age": 22
		}
	`

	var result Student
	
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
 
	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}