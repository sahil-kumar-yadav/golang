package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time pacakage in go")
	// Get current time
	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)

	// Format time
	formattedTime := currentTime.Format("2006-01-02 15:04:05") // "2006-01-02 15:04:05 yeh fixed hai
	fmt.Println("Formatted time: ", formattedTime)
	formattedTime = currentTime.Format("2006-01-02")
	fmt.Println("Formatted time: ", formattedTime)
	formattedTime = currentTime.Format("15:04:05")
	fmt.Println("Formatted time: ", formattedTime)

	formattedTime = currentTime.Format("15:04:05, Monday")
	fmt.Println("Formatted time: ", formattedTime)

	formattedTime = currentTime.Format("03:04 PM") // PM he dalna padega
	fmt.Println("Formatted time: ", formattedTime)

	// Parse time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2022-05-15 12:34:56") // layoutstring , formatstring
	if err != nil {
		fmt.Println("Error parsing time: ", err)
	} else {
		fmt.Println("Parsed time: ", parsedTime)
	}
	fmt.Printf("--------------------------------\n")

	// getting date , day , month


	// Time zone
	currentTime = time.Now().In(time.FixedZone("PST", -8*60*60)) // -8 hours
    fmt.Println("Current time in PST: ", currentTime)

    // Duration
    duration := time.Minute * 5
    fmt.Println("Duration: ", duration)

    // Add duration to current time
    newTime := currentTime.Add(duration)
    fmt.Println("New time after adding duration: ", newTime)

    // Subtract duration from current time
    newTime = currentTime.Add(-duration)
    fmt.Println("New time after subtracting duration: ", newTime)

    // Sleep for a duration
    time.Sleep(time.Second * 5)
    fmt.Println("Sleeping for 5 seconds...")

}
