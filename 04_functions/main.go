package main

import (
	"fmt"
)

func simplefunction() {
	fmt.Println("Simple Function Called...")
}

func add(a, b int) int {
	return a + b
}

func add1(a, b int) (result int) { // curly brackets yahi se start karna padega
	result = a + b
	return
}

func newfunction(a float64, b float64) (result float64) { // curly brackets yahi se start karna padega
	result = a + b
	return
}


func main() {
	fmt.Println("Function in Go language")
	simplefunction()
	result := add(5, 7)
	result1 := add1(5, 7)
	result2 := newfunction(5, 7)


	fmt.Println("Result:", result)
	fmt.Println("Result1:", result1)
	fmt.Println("Result1:", result2)


}
