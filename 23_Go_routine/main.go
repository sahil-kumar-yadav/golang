package main

import (
	"fmt"
	"time"
)

func firstFunction() {
	fmt.Println("firstFunction Started")
	// do some task
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("firstFunction Completed successfully")
	fmt.Println()
}

func secondFunction() {
	fmt.Println("secondFunction Started")
	// do some task
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("secondFunction Completed successfully")
	fmt.Println()
}

func main() {
	fmt.Println("Go Routines in GoLang ................")
	fmt.Println()

	// case 1: simple function run
	// firstFunction() // 1
	// secondFunction() // 2 execute

	// Case 2: goroutine in first function
	// go firstFunction() // this will not be executed because until the first function executes the main function will be executed
	// secondFunction() // this will executed only

	// Case 3: goroutine in both function without waiting in main function
	// go firstFunction() // not executed
	// go secondFunction() // not executed


	// case 4: goroutine in both function with waiting time in main function greater than first function
	// first --> 1000
	// secondFunction --> 2000
	// go firstFunction()  // executed because waiting time is 1000
	// go secondFunction() // not executed fully
	// time.Sleep(1100 * time.Millisecond) // this will help in this case

	// case 5:goroutine in both function with waiting time greater than timeout of both goroutines
	// first --> 1000
	// secondFunction --> 2000
	go firstFunction()  // executed because waiting time is 2100
	go secondFunction() // executed because waiting time is 2100
	// waiting for goroutine to complete
	time.Sleep(2100 * time.Millisecond) // this will help in this case

	
	fmt.Println("Main function Completed successfully")
}
