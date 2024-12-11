package main

import (
    "fmt"
    
)

func main() {
	fmt.Println("Arrays in Golang")

	// golang by default intializes arrays with their default values
	// for interger --> 0
	// for boolean --> false
	// for string --> ""
	// for float --> 0
	// for pointer --> nil

    var numbers [5]int
    fmt.Println("Arrays contentes are",numbers)

    // initialize array with empty strings
    var emptyStrings = [3]string {}

    fmt.Println("Arrays contentes are",emptyStrings)
    fmt.Printf("Arrays contentes are %q \n",emptyStrings) //%q is quoted string


	

	var name[5]string
	
	name[0] = "Alice"
	name[3] = "Bob"

	fmt.Println("Arrays contentes are",name)

	// intialize arrays
	var numArrays = [5] int  {1,2,3,4,5}

	fmt.Println("numarrays is",numArrays)
	fmt.Println("length of numarray is",len(numArrays))

	fmt.Println("number at 2nd index of numarray is",numArrays[2])
}