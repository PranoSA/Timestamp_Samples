package main

import (
	"fmt"
	"time"
)

func main() {
    // Define a location (UTC in this case, since leap seconds are universal)
    location, err := time.LoadLocation("UTC")
    if err != nil {
        fmt.Println("Error loading location:", err)
        return
    }

    // Define dates around a simulated leap second event
    // For example, adding a leap second between 23:59:59 and 00:00:00 of the next day
    dateBeforeLeap := time.Date(2021, time.December, 31, 23, 59, 59, 0, location)
    dateAfterLeap := time.Date(2022, time.January, 1, 0, 0, 0, 0, location)

    // Calculate the difference in seconds
    duration := dateAfterLeap.Sub(dateBeforeLeap)
    seconds := duration.Seconds()

    // Print the results
    fmt.Printf("Date before leap second: %v\n", dateBeforeLeap)
    fmt.Printf("Date after leap second: %v\n", dateAfterLeap)
    fmt.Printf("Seconds between the two dates: %v\n", seconds)
}
