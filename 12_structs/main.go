package main

import (
	"fmt"
)

type employe struct {
	FirstName string
	LastName  string
	age       int
	email     string
}

func main() {
	fmt.Println("Structs in Golang")
	var e1 employe
	fmt.Println("employe e1 is ", e1)
	e1.FirstName = "John"
	e1.LastName = "Doe"
	e1.age = 20
	e1.email = "johndoe@example.com"
	println("----------------")
	fmt.Println("employe e1 is ", e1)
	fmt.Println("Full Name : ", e1.FirstName+" "+e1.LastName)
	fmt.Println("Age : ", e1.age)
	fmt.Println("Email : ", e1.email)

	// 2nd method to intialize and create a stuct data
	fmt.Println()
	e2 := employe{
		FirstName: "kumar",
		LastName:  "vishwas",
		email:     e1.email,
		age:       24}

	fmt.Println("employe e2 data is ", e2)

	// using new keyword
	fmt.Println()
	var e3 = new(employe)
	e3.FirstName = "Ravi"
	e3.LastName = "Singh"
	e3.age = 25
	e3.email = "ravisingh@example.com"
	fmt.Println("employe e3 data is ", e3)
	fmt.Println("employe e3 data is ", *e3)


}
