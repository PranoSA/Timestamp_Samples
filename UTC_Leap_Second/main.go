package main

import (
	"fmt"
	"time"
)

func main() {
    // Define the date and time
    dateTime := time.Date(2015, time.June, 30, 12, 0, 0, 0, time.UTC)
    
    // Convert to Unix timestamp (seconds since epoch)
    unixTimestamp := dateTime.Unix()

    // Print the results
    fmt.Printf("Date and Time: %v\n", dateTime)
    fmt.Printf("Unix Timestamp: %v\n", unixTimestamp)
}