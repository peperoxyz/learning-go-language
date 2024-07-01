package main

import (
	"encoding/json"
	"fmt"
)

// JSON Decoding to Slice of Struct
// Struct student
type Student struct {
	FullName string	`json:"full_name"`
	Email string 	`json:"email"`
	Age int			`json:"age"`
}

func main () {
	var jsonString = `
	[
		{
			"full_name": "Dea Ananda",
			"email": "dea@gmail.com",
			"age": 22
		},
		{
			"full_name": "Bru Xhaol",
			"email": "bruxhaol@gmail.com",
			"age": 28
		}
	]
	`

	var result []Student
	
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, val := range result {
		fmt.Printf("Index %d: %+v\n", i+1, val )
	}

 

}