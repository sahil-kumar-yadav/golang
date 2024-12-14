package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Advance file handling")

	// reading file through buffer
	file, err := os.Open("./18_File_handling_1/sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // close the file when we're done

	// create a buffer of size 1000
	buffer := make([]byte, 1000)

	// read file into buffer
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}

	// print the content of the buffer
	fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	fmt.Println()

}
