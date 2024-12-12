package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch statements")

	temprature := 10
	// we can use expressions in case statements also
	switch {
		case temprature > 25:
            fmt.Println("It's a hot day")
        case temprature >= 15 && temprature <= 25:
            fmt.Println("It's a warm day")
        default:
            fmt.Println("It's a cold day")
    }
	fmt.Printf("\n")


	day := time.Now().Weekday()
	fmt.Println("day: ", day)
	fmt.Printf("type of data is %T\n",day)

	fmt.Printf("--------------------------------\n")
	switch day {
	case time.Monday:
		fmt.Println("Today is Monday")
    case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
    case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
    case time.Saturday:
		fmt.Println("Today is Saturday")
	default:
		fmt.Println("Hurray !! Enjoy Today is Sunday")
	}
}
