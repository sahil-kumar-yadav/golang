package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File handling in Go lang")

	// creating a new file
	file, err := os.Create("newfile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // close the file when we're done

	// creating a new file inside a directory
	sampleFile, error := os.Create("./18_File_handling/sample.txt")
	if error != nil {
		fmt.Println("Error creating file:", error)
		return
	}
	defer sampleFile.Close() // close the file when we're done

	// writing to a file
	data := []byte("Hello, World!\n")
	n, err := file.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", n)
	if err != nil {
		fmt.Printf("Error writing to file: %v, wrote %d bytes\n", err, n)
		return
	}

	// reading from a file
	content, err := os.ReadFile("newfile.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File content: %s\n", content)

	// updating a file

	// file handling program
	// opening a file
	// checking if a file exists
	// creating a file
	// truncating a file
	// closing a file
	// reading from a file
	// writing to a file
	// seeking in a file
	// closing a file
	// error handling
	// file permissions
	// file stat
	// file info
	// reading a file
	// writing to a file
	// closing a file
	// error handling
	// file permissions
	// file stat
	// file info
	// reading a file
	// writing to a file
	// closing a file
	// error handling
	// file permissions
	// file stat
	// file info
	// reading a file
	// writing to a file
	// closing a file
	// error handling
	// file permissions
	// file stat
	// file info

	// reading a file
	// writing to a file
	// closing a file
	// error handling
	// file permissions
	// file stat
	// file info
	// directory handling
	// reading directory contents
	// creating a directory
	// removing a directory
	// renaming a file or directory
	// copying a file or directory
	// moving a file or directory
	// checking if a file or directory exists
	// checking if a file is a directory
	// checking if a file is a symbolic link
	// checking if a file is a block device
	// checking if a file is a character device
	// checking if a file is a pipe or FIFO
	// checking if a file is a socket
	// checking if a file is a symbolic link to a directory
	// checking if a file is a symbolic link to a file

}
