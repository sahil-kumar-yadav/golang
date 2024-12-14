package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Data Conversion in Go")

	var num int = 12
	// fmt.Printf("number is %f\n",num)
	fmt.Printf("number is %d\n",num)
	fmt.Printf("Type of number is %T\n",num)

	// converting number to float64

	var floatnum float64 = float64(num)
	floatnum = floatnum + 1.23
	fmt.Printf("floatnum is %f\n", floatnum)
	fmt.Printf("Type of floatnum is %T\n",floatnum)

	// converting integer to string
	println("-------------------")

	num = 123
	var str string = strconv.Itoa(num) // interger to alphabetic
	fmt.Printf("str is %s\n", str)
	fmt.Printf("Type of str is %T\n", str)

	// converting string to integer
	fmt.Println()
	stringValue := "123"
	intValue ,_ := strconv.Atoi(stringValue)
	fmt.Printf("intValue is %d\n", intValue)
	fmt.Printf("Type of intValue is %T\n", intValue)

	fmt.Println("-------------------")
	
	// converting string to float64

	stringValue = "12.34"
	floatValue, _ := strconv.ParseFloat(stringValue, 64) // (value,bit)
	fmt.Printf("floatValue is %f\n", floatValue)
	fmt.Printf("Type of floatValue is %T\n", floatValue)

}