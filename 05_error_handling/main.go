package main

import (
	"fmt"
)

func divide(a, b float64) float64 {
	return a / b
}

// error apne aap me ek type hai go lang ke ander
// error handling
func divide1(a, b float64)( float64,error) {

	if(b == 0){
		return 0, fmt.Errorf("cannot <-- Cannot divide by zero")  // return error message and nil value
        // or return 0, errors.New("Cannot divide by zero")  // return error message and error object
	}
	return a / b , nil // return no error
}



func main() {
	fmt.Println("Error handling")
	ans := divide(10,4)
	ans1 := divide(10,0)
	fmt.Println("Answer:", ans)
	fmt.Println("Answer:", ans1) // it will give +infinity
	fmt.Println("now with error message")
	// ans2, _ := divide1(10,0) // yaha hum error use nahi karna chate , so underscore lagaya
	ans2, err := divide1(10,0)

	
	if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Answer:", ans2)
    }

	// underscore hum tabh use karte hai jab hum function ka ek output ignore karna chate hai
}
