package main

import (
	"fmt"
	"time"
)

/**
 *
 *
 * We Will Want to Illustrate some Time zone Caveats and how to handle them
 *
 * For example -> When American Samoa changed from UTC-11 to UTC-10 on 2011-09-24
 *
 * We Want to See that the Go Time Package will handle this change
 */

//Mon Jan 2 15:04:05 MST 2006 is the reference time for the Go Time Package

var american_samoa_Time_2 = "2011-09-24T00:00:00-1000"

// Here We Should Expect the Time to be 1 Second Later"

var iso_layout = "2006-01-02T15:04:05-0700"

var custom_random_format = "2006Year 01Month 02Day 15Hour 04Minute 05Second -07:00GMT"

var american_samoa_Time_1 = "2011Year 09Month 23Day 23Hour 59Minute 59Second -11:00GMT"


func main() {


	//parse american_samoa_Time_1

	t, err := time.Parse(custom_random_format, american_samoa_Time_1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)

	// convert to unix timestamp
	unix_time := t.Unix()

	fmt.Println(unix_time)
	

}