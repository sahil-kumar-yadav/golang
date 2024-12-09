package main

import "fmt"


// same rule applies to functions

func Publicfunction(){
	fmt.Println("This is a Public function")
}
func privatefunction(){
	fmt.Println("This is a Private function")
}
func main() {
	fmt.Println("Datatype and veriable")
	var name string = "John Doe"
	fmt.Println(name)
	var version = "latest version" // here we use varible with or without datatype and ver
	fmt.Println(version) 
	var versionnumber = 12
	fmt.Println(versionnumber)
	
	var money int = 123
	var currency float64 = 12.12
	fmt.Println(money)
	fmt.Println(currency)

	var country string = "India"
	var ruppie float64 = 100
	fmt.Printf("The currency of %s is %f\n", country, ruppie)

	var flag bool = false
	fmt.Println("apple is fruit ",flag)

	const pi = 3.14 // value can't be changed
	fmt.Println("Value of pi is ", pi)

	// shortcut for variable declaration
	personname := "kunal"
	fmt.Println("Person's name is ", personname)

	var Public = "data is public" // this variable can be used by outside packages (first letter capital) // import export
	var private = "data is private" // this variable can't be used outside of this file
	fmt.Println(Public)
	fmt.Println(private)

	
}
