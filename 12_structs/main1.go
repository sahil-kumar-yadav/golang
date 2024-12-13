package main

import (
	"fmt"
)

type BasicDetails struct {
	name string
	Age  int
}

type Contact struct {
	phone string
	email string
}

type Address struct {
	city  string
	state string
	zip   string
}

type Employe struct {
	employeDetail BasicDetails
	contact       Contact
	address       Address
}

func main() {
	fmt.Println("More in Structures")
	var e1 Employe
	e1.employeDetail = BasicDetails{
		name: "John Doe",
		Age:  30,
	}
	e1.contact = Contact{
		phone: "+1-123-456-7890",
		email: "johndoe@example.com",
	}
	e1.address = Address{
		city:  "New York",
		state: "NY",
		zip:   "10001",
	}
	fmt.Println("Employe e1 is ", e1)
	fmt.Println()
	fmt.Println("Full Name : ", e1.employeDetail.name)
	fmt.Println("Age : ", e1.employeDetail.Age)
	fmt.Println("Phone : ", e1.contact)
	fmt.Println("Email : ", e1.contact.email)
}
