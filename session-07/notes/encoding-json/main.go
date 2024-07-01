package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:name`
	Age      int    `json:age`
}

func main() {
	var object = []User{{"Dea", 22}, {"Ananda", 21}}

	jsonData, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)
}