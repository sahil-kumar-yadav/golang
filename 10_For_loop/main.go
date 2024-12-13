package main

import (
	"fmt"
)

// Go lang me sirf ek he loop hai for loop
func main() {
	fmt.Println("For Loop in GOlang")

	fmt.Println("Printing numbers from 0 to 5")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// infinite loop

	counter := 1

	for {
		fmt.Println("Infinite loop counter:", counter)
		counter++
		if counter == 5 {
			break
		}
	}
	fmt.Printf("--------------------------------\n")
	// range keyword
	numbers := []int{12, 29, 13, 49, 35}
	fmt.Println("Printing numbers using range keyword")

	for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

	fmt.Printf("------------------------")
	// iterating over string 
	data := "This is simple string"

	for index, value := range data {
		fmt.Printf("Index: %d, Value: %c\n", index, value)
	}

    fmt.Printf("------------------------\n")
    stringdata := []string{"This","is","a","string","array"}

	for index, value := range stringdata {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

}
