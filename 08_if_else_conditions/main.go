package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else condition \n") // here also \n works

	var age int = 10

	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// if elseif statement

	var number int = 4

	// syntax essa he likna padega , brackets ke baad else if aayega
	if number > 0 && number < 5 {
		fmt.Println("Number is between 0 and 5")
	} else if number > 0 && number < 10 {
		fmt.Println("Number is between 5 and 10")
	} else if number > 0 && number > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is 0 or negative")
	}

}
