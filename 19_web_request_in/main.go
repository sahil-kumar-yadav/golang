package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web request in Go lang")
	// https://jsonplaceholder.typicode.com/todos/1

	// Geting data from http://jsonplaceholder.typicode.com
	// and printing it in JSON format
	// Example:
	// {
	//   "userId": 1,
	//   "id": 1,
	//   "title": "delectus aut autem",
	//   "completed": false
	// }

	// TODO: Implement HTTP GET request to fetch data from the specified URL and print it in JSON format.
	// HINT: You can use the "net/http" package in Go to make HTTP requests.

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()
	// fmt.Println(res)
	fmt.Printf("Data Type of response: %T \n", res) // *http.Response
	bodyBytes, err := io.ReadAll(res.Body)
	if err!= nil {
        fmt.Println("Error reading response body:", err)
        return
    }
	fmt.Println("Response Body: ",bodyBytes) // response body in bytes
	fmt.Println(string(bodyBytes)) // response body in string

}
