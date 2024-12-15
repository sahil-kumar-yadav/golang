package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// {
//   "userId": 1,
//   "id": 1,
//   "title": "delectus aut autem",
//   "completed": false
// }

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func makingGetRequest() {
	url := "https://jsonplaceholder.typicode.com"
	// url := "https://jsonplaceholder.typicode.commm"
	res, err := http.Get(url + "/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	// checking the response
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d - %s\n", res.StatusCode, http.StatusText(res.StatusCode))
		return
	}

	// reading the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// fmt.Println("Response Body: ", bodyBytes) // response body in bytes
	fmt.Println(string(bodyBytes)) // response body in string

	fmt.Println("--------------------------------")

	// unmarshal response body
	var todo Todo
	err = json.Unmarshal(bodyBytes, &todo)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	fmt.Printf("Todo: %+v\n", todo)

	fmt.Println("--------------------------------")
	// genrally used methods -> decoder method
	// to directly store data in our struct
	newRes, err := http.Get(url + "/todos/21")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	var todoA Todo
	err = json.NewDecoder(newRes.Body).Decode(&todoA)
	if err != nil {
		fmt.Println("Error Decoding in todo1:", err)
		return
	}
	fmt.Println("TodoA: ", todoA)

	fmt.Println("--------------------------------")
}

func makingPostRequest() {
	postUrl := "https://jsonplaceholder.typicode.com/todos"

	// creating a new todo
	newTodo := Todo{
		UserId:    1,
		Id:        21,
		Title:     "new todo from go",
		Completed: true,
	}
	// conveting struct to json
	newTodoJson, err := json.Marshal(newTodo)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	fmt.Println("New Todo JSON: ", string(newTodoJson))
	// convert to json string
	jsonString := string(newTodoJson)
	// convert string data to reader
	reader := strings.NewReader(jsonString)

	// making a POST request
	res, err := http.Post(postUrl, "application/json", reader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("res body data:", string(data))
	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Println("--------------------------------")
	// res, err = http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", bytes.NewReader(newTodoJson))

}
func main() {
	fmt.Println("CRUD API in GOlang")
	// https://jsonplaceholder.typicode.com/todos/1
	// makingGetRequest()
	makingPostRequest()

}
