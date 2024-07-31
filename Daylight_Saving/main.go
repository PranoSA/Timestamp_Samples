package main

import (
	"fmt"
	"time"
)

func main() {
    // Define the location (e.g., "America/New_York")
    location, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println("Error loading location:", err)
        return
    }

    // Define dates around the DST transition (for demonstration, we'll use March 14, 2021)
    // On March 14, 2021, DST starts at 2:00 AM, which means the clock jumps from 1:59:59 AM to 3:00:00 AM
    dateBeforeDST := time.Date(2021, time.March, 14, 1, 59, 59,0, location)
    dateAfterDST := time.Date(2021, time.March, 14, 3, 0, 0, 0, location)

    // Calculate the difference in seconds
    duration := dateAfterDST.Sub(dateBeforeDST)
    seconds := duration.Seconds()

    // Print the results
    fmt.Printf("Date before DST: %v\n", dateBeforeDST)
    fmt.Printf("Date after DST: %v\n", dateAfterDST)
    fmt.Printf("Seconds between the two dates: %v\n", seconds)
}