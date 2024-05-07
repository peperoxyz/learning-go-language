package main

import (
	"fmt"
	"strings"
)

func main() {
	// #1 MAP: Tipe data asosiatif yang berbentuk key-value pair (pasangan keyword-value)

	// first way of declaring map
	var student map[string]string // declaring map
	student = map[string]string{} // init

	// inserting value
	student["name"] = "Dea"
	student["student_id"] = "3312001121"
	student["major"] = "Informatics Engineering"

	fmt.Println(student)

	// second way
	var assets = map[string]int {
		"houses" : 2,
		"cars" : 2,
		"motor cycle" : 2,
	}

	fmt.Println("Dea's Houses: ", assets["houses"] )
	fmt.Println()

	// #2 iterasi map dengan for = range looping
	var chicken map[string]int
	chicken = map[string]int{
		"January" : 30,
		"February" : 10,
		"March" : 5,
		"April" : 15,
	}

	for key, val := range chicken {
		fmt.Println(key, ":", val)
	}
	
	// #3 menghapus item pada map dengan fungsi delete(map, "key")
	delete(chicken, "January")
	fmt.Println("Setelah menghapus january: ", chicken)

	// #4 deteksi keberadaan item dengan key tertentu
	var song = map[string]string {
		"singer" : "Nadine Amizah",
		"title" : "Rayuan Perempuan Gila",
		"release_date" : "2023",
	}

	var value, isExist = song["title"]

	if isExist {
		fmt.Println("The value is exist:", value)
	} else {
		fmt.Println("Item doesn't exist")
	}

	// #5 Kombinasi Slice and Map: slice yang isinya berupa kumpulan map
	var songs = []map[string]string {
		map[string]string {"title": "Mendarah", "singer": "Nadin Amizah", "release_date": "2023"},
		map[string]string {"title": "Puisi", "singer": "Jikustik", "release_date": "2006"},
	}

	fmt.Println(strings.Repeat("==", 30))

	// for key, val := range songs {
	// 	fmt.Println(key, ":", val)
	// }

	for _, songs := range songs {
		fmt.Println(songs["title"], songs["singer"])
	}



}