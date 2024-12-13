package main

import (
	"fmt"
)

// pointer function
func changeValuePointer(num *int) {
    *num = *num + 10
}

func changeValuePointerReturn(num *int) *int {
    *num = *num + 10
    return num
}

func main() {
	fmt.Println("Pointers in Go lang")

	num := 2
	var ptr *int = &num
	fmt.Printf("type of ptr is %T\n", ptr)
	fmt.Println("The value at ptr is:", *ptr)
	fmt.Println("value in pointer ", ptr)
	fmt.Println("--------------------------------")
	// another simple way

	num1 := 12
	ptr1 := &num1

	fmt.Println("The value at ptr1 is:", *ptr1)
	fmt.Println("Value in pointer1 ", ptr1)
	*ptr1 = *ptr1 + num1
	fmt.Println("The value at ptr1 now is:", *ptr1)

	// nil pointer
	var nilPointer *int = nil
    if nilPointer == nil {
        fmt.Println("nil pointer")
    }
    fmt.Println("--------------------------------")
    // pointer to a pointer
    var ptr2 **int = &ptr1
    fmt.Println("The value at ptr2 is:", **ptr2)
    fmt.Println("Value in pointer2 ", ptr2)
    **ptr2 = **ptr2 + *ptr1
    fmt.Println("The value at ptr2 now is:", **ptr2)

    fmt.Println("--------------------------------")
    // pointer to a slice
    var slice []int = []int{1, 2, 3, 4, 5}
    var slicePtr *[]int = &slice
	sclicePtr2 := &slice
    fmt.Println("The value at slicePtr is:", *slicePtr)
    fmt.Println("Value in slicePtr ", slicePtr)
    fmt.Println("Value in slicePtr2 ", sclicePtr2)

	fmt.Println("--------------------------------")
    // pointer to a function
    fmt.Println("Before function call")
    fmt.Println("The value at ptr1 is:", *ptr1)
    fmt.Println("The value at ptr2 is:", **ptr2)
    fmt.Println("--------------------------------")
    // passing pointer to a function
    changeValuePointerReturn(ptr1)
    fmt.Println("After function call")
    fmt.Println("The value at ptr1 after function call is:", *ptr1)
    fmt.Println("The value at ptr2 after function call is:", **ptr2)
	changeValuePointer(ptr1)
    fmt.Println("The value at ptr1 after function call is:", *ptr1)
    *ptr2 = changeValuePointerReturn(*ptr2)
    fmt.Println("The value at ptr1 after function call is:", *ptr1)



}
