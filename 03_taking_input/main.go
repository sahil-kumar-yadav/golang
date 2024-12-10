package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Your Name")
	// var name string
	// fmt.Scan(&name) // ye sirf ek word lega
	// fmt.Println("Hello,", name)

	// taking input with spaces
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)
}
