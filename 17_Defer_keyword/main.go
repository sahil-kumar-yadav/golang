package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer Keyword in Go lang")

	// defer keyword wali line last me exicute hogi
	// defer function execution order is LIFO (Last In First Out)

    // Why to use defer keyword?
    // Deferred functions are useful when you want to ensure that a function executes after the surrounding function returns.
    // They can be used for cleanup operations, logging, or other tasks that should be performed regardless of whether the surrounding function exits normally or with an error.
    // Deferred functions can also be used to manage resources, such as file handles or network connections, by ensuring they are properly closed or released when they are no longer needed.
    // Additionally, deferred functions can be used to implement error handling, where you can perform cleanup actions even if an error occurs within the surrounding function.
    // Deferred functions can also be used to implement control flow, such as breaking out of nested loops or returning early from a function.

   
	// Defer keyword is used to defer the execution of a function until the surrounding function returns.
	// The deferred function's arguments are evaluated when the surrounding function returns, not when the defer statement is executed.

	// Example 1: Defer function execution
	fmt.Println("starting line")
	defer fmt.Println("Middle line")
	fmt.Println("Last line")

	// Example 2: Multiple deferred functions
	fmt.Println("\nExample 2:")
	defer fmt.Println("Deferred function 1")
	fmt.Println("Deferred function 2")
	defer fmt.Println("Deferred function 3")
	fmt.Println("Surrounding function execution")

}
