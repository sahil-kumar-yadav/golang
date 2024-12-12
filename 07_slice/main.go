package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices for dynamic arrays")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)
	fmt.Println("length of original slice :", len(numbers))
	numbers = append(numbers, 6,7)
	fmt.Println("After appending:", numbers)
	fmt.Println("length of new slice :", len(numbers))

	fmt.Printf("Type of number : %T\n", numbers)

	// empty slice
	name := []string{}
	fmt.Println("Empty slice:", name)

	// lenght and capacity

	num := []int{1,2}
	fmt.Println("length of num slice :", len(num))
	fmt.Println("capacity of num slice :", cap(num))
	num = append(num,3,4,5,6,7)
	fmt.Println("After appending length :", len(num))
	fmt.Println("After appending capacity of num slice :", cap(num))

	fmt.Println("-------------------------------------------------------")
	// slice using make function

	var length int = 3
	var capacity int = 5
	values := make([]int,length,capacity)
	fmt.Println("Values slice:", values)
	fmt.Println("length of values slice :", len(values))
	fmt.Println("capacity of values slice :", cap(values))

	fmt.Printf("\n")
	values = append(values, 3,4)
	fmt.Println("Values slice:", values)
	fmt.Println("length of values slice :", len(values))
	fmt.Println("capacity of values slice :", cap(values))


	fmt.Printf("\n")
	values = append(values, 5)
	fmt.Println("Values slice:", values)
	fmt.Println("length of values slice :", len(values))
	fmt.Println("capacity of values slice :", cap(values)) // capacity of values double kar deta hai full hone pe


	fmt.Printf("\n")
	// empty slice using make function
	empty := make([]int, 0)
    fmt.Println("Empty slice:", empty)
    fmt.Println("length of empty slice :", len(empty))
    fmt.Println("capacity of empty slice :", cap(empty))
	
	fmt.Printf("\n")
    empty = append(empty, 1,2,3)
    fmt.Println("Empty slice:", empty)
    fmt.Println("length of empty slice :", len(empty))
    fmt.Println("capacity of empty slice :", cap(empty))
}
