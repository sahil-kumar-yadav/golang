package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the waitgroup counter when the goroutine completes
	fmt.Printf("Started worker %d\n", i)
	// do some work here
	//...
	fmt.Printf("Completed worker %d\n", i)

}

func main() {
	// fmt.Println("Sync waitgroup in Go")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // add sync waitgroup
		go worker(i, &wg)
	}

	// wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Main function completed")
}
