package main

import (
    "fmt"
)

func main() {
    // fmt.Println("Hello, World!")
	name := "Alice"
	age := 25
	height := 180.6

	fmt.Println("nama:", name,"age:", age, "height:", height)

	// printf  (print formatted)

	fmt.Printf("name is %s\n",name) // \n print in new line
	fmt.Printf("age is %d\n",age)
	// fmt.Printf("height is %d cm\n",height) // this will give error
	fmt.Printf("height is %.2f cm\n",height) // .2f means point ke baad 2 values

	// to print the type of data type we use %T

	fmt.Printf("Type of name is %T\n",name)
	fmt.Printf("Type of age is %T\n",age)
	fmt.Printf("Type of height is %T\n",height)

	// %s, %d , %f are format specifiers

	fmt.Printf("name is %s , age is %d , height is %f\n",name,age,height)
}