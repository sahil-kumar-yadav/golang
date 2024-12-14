package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings in Go")

	// split function
	data := "apple,banana,mango,orange" // csv comma separated values
	fmt.Println("data: ", data)
	fruits := strings.Split(data, ",") // second parameter is comma
	fmt.Println("Fruits:", fruits)
	fmt.Println("data: ", data)
	data = "apple.banana.mango.orange"
	fruits = strings.Split(data, ".")
	fmt.Println("Fruits:", fruits)

	fmt.Printf("Type of fruits is %T \n", fruits) // string array or string slice

	// count function

	data = "This is the , simpe data and this is used to count the,number of the in this data"
	fmt.Println("data: ", data)
	fmt.Println("Number of the in data is ", strings.Count(data, "the"))

	// trim function // starting or ending se space clear karta hai
	data = "  This is a   string  "
	fmt.Println("data: ", data)
	fmt.Println("Trimmed data: ", strings.TrimSpace(data))

	// replace function
	data = "This is a simple string"
	fmt.Println("data: ", data)
	fmt.Println("Replaced data: ", strings.Replace(data, "simple", "complex", -1)) // -1 means replace all occurrences

	// Join the data
	str1 := "hello"
	str2 := "world"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("result: ", result)
	fmt.Println()
	newdata := []string{"This", "is", "a", "string", "array"}
	fmt.Println("data: ", newdata)
	fmt.Println("Joined data: ", strings.Join(newdata, " ")) // join data with spaces seperator
	fmt.Println()

	// Contains function
	data = "This is a simple string"
	fmt.Println("data: ", data)
	fmt.Println("Contains 'simple': ", strings.Contains(data, "simple"))

	// Index function
	fmt.Println()
	data = "This is a simple string"
	fmt.Println("data: ", data)
	fmt.Println("Index of 'simple': ", strings.Index(data, "simple"))

	// LastIndex function
	fmt.Println()
	data = "This is a simple string"
	fmt.Println("data: ", data)
	fmt.Println("Last index of 'simple': ", strings.LastIndex(data, "simple"))

	// ToLower function
	data = "This IS A SIMPLE STRING"
	fmt.Println("data: ", data)
	fmt.Println("To lower case: ", strings.ToLower(data))

	// ToUpper function
	data = "this is a simple string"
	fmt.Println("data: ", data)
	fmt.Println("To upper case: ", strings.ToUpper(data))

	// Repeat function
	data = "hello"
	fmt.Println("data: ", data)
	fmt.Println("Repeat data: ", strings.Repeat(data, 3))

	// ReplaceRune function
	// data = "This is a simple string"
	// fmt.Println("data: ", data)
	// fmt.Println("Replace 'a' with '*': ", strings.ReplaceRune(data, 'a', '*', -1)) // -1 means replace all occurrences

	// other sting functions
	// ContainsAny function
	// Fields function
	// HasPrefix function
	// HasSuffix function
	// Map function
	// NewReplacer function
	// ReplaceAll function
	// ToTitle function
	// ToLowerSpecial function
	// ToUpperSpecial function
	// ToCamelCase function
	// ToSnakeCase function
	// ToPascalCase function
	// ToKebabCase function
	// ToUnderscoreCase function
	// ToScreamingSnakeCase function
	// ToConstantCase function
	// ToEnumCase function
	// ToGoName function
	// ToSnakeName function
	// ToCamelName function
	// ToPascalName function
	// ToKebabName function
	// ToUnderscoreName function
	// ToScreamingSnakeName function
	// ToConstantName function
	// ToEnumName function

}
