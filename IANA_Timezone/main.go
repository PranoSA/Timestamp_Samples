package main

import (
	"fmt"
	"time"
)

func main() {
    // Load the location for Pacific Time (Los Angeles)
    location, err := time.LoadLocation("America/Los_Angeles")
    if err != nil {
        fmt.Println("Error loading location:", err)
        return
    }

    // Define a date in PST (Standard Time)
    datePST := time.Date(2024, time.January, 15, 12, 0, 0, 0, location)
    fmt.Printf("PST Date: %v\n", datePST)

    // Define a date in PDT (Daylight Time)
    datePDT := time.Date(2024, time.July, 15, 12, 0, 0, 0, location)
    fmt.Printf("PDT Date: %v\n", datePDT)

    // Print the time zone names and offsets
    _, offsetPST := datePST.Zone()
    _, offsetPDT := datePDT.Zone()
    fmt.Printf("PST Offset: %d seconds (%v)\n", offsetPST, time.Duration(offsetPST)*time.Second)
    fmt.Printf("PDT Offset: %d seconds (%v)\n", offsetPDT, time.Duration(offsetPDT)*time.Second)

	//try phoenix now, which doesn't have Daylight Savings Time
	location, err = time.LoadLocation("America/Phoenix")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Define a date in Phoenix
	datePhoenix := time.Date(2024, time.January, 15, 12, 0, 0, 0, location)
	fmt.Printf("Phoenix Date: %v\n", datePhoenix)
	
	// Print the time zone names and offsets
	_, offsetPhoenix := datePhoenix.Zone()
	fmt.Printf("Phoenix Offset: %d seconds (%v)\n", offsetPhoenix, time.Duration(offsetPhoenix)*time.Second)
	
	//try June to demonstrate no DST
	datePhoenix = time.Date(2024, time.June, 15, 12, 0, 0, 0, location)
	fmt.Printf("Phoenix Date: %v\n", datePhoenix)

	// Print the time zone names and offsets
	_, offsetPhoenix = datePhoenix.Zone()
	fmt.Printf("Phoenix Offset: %d seconds (%v)\n", offsetPhoenix, time.Duration(offsetPhoenix)*time.Second)

	// now -> current location and time
	now := time.Now()

	// Print the current time
	fmt.Printf("Current Time: %v\n", now)

	// Get the time zone and offset for the current location
	zone, offset := now.Zone()
	fmt.Printf("Time Zone: %v\n", zone)

	// Print the offset in seconds
	fmt.Printf("Offset: %d seconds (%v)\n", offset, time.Duration(offset)*time.Second)

	//convert to unix timestamp
	unixTimestamp := now.Unix()
	fmt.Printf("Unix Timestamp: %v\n", unixTimestamp)

// unix 64?
	unixTimestamp64 := now.UnixNano()
	fmt.Printf("Unix Timestamp 64: %v\n", unixTimestamp64)
	
}