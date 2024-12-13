package main

import (
    "fmt"
)

func main() {
    fmt.Println("Maps in GoLang")

	// Here map is Unordered and dynamic
	// map of name <-> grades
	// map[key]value

	// map creation 
	person := map[string]int{
		"Ac": 85,
        "Dc": 90,
        "ADC": 75,
	}
	
	fmt.Println("person data is  ", person)
	fmt.Printf("\n")

	// map using make function
	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 85
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 75
	fmt.Println("Student grades:", studentGrades)
	studentGrades["Charlie"] = 95
	fmt.Println("Student grades:", studentGrades)
	
	// value using key
	fmt.Println("Student grade for Alice:", studentGrades["Alice"])
	fmt.Println("Student grade for Raushan:", studentGrades["Raushan"]) // not presesnt so it will show zero
	fmt.Printf("\n")
	// delete key-value pair
	delete(studentGrades,"Bob")
	fmt.Println("Student grades:", studentGrades)

	// check if a key-value pair exists in map or not

	grades, exists := studentGrades["prabhat"]
	fmt.Println("Student prabhat exists:",exists)
	fmt.Println("Student prabhat's grade:", grades)
	fmt.Printf("\n")
	grades1, exists1 := studentGrades["Alice"]
	fmt.Println("Student Alice exists:",exists1)
	fmt.Println("Student Alice's grade:", grades1)

	fmt.Printf("--------------------------------\n")

	employdata := make(map[string]int) // name -> id
	employdata["John Doe"] = 123
	employdata["Alice"] = 21
	employdata["Bob"] = 321
	employdata["Doe"] = 33
	employdata["Jhonny"] = 88

	for k,v := range employdata{
		fmt.Printf("Name: %s, ID: %d\n", k,v)
	}

	
}