package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	IsMen bool   `json:"ismen"`
}

func main() {
	fmt.Println("Json in GO Lang")
	p1 := Person{Name: "John Doe", Age: 20, IsMen: true}
	fmt.Printf("Type of person is %T\n", p1)
	fmt.Println("Person is :", p1)
	// object se json me convert
	// converting into json --> marshalling the struct to json
	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	fmt.Printf("Type of jsonData is %T\n", jsonData)
	fmt.Println("Json data is :", string(jsonData))

	// converting from json --> unmarshalling the json to struct
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	fmt.Printf("Type of p2 is %T\n", p2)
	fmt.Println("Person from json is :", p2)

}
