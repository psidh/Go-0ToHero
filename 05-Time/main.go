package main

import (
	"fmt"
	"time" // Rename the standard library's time package
)

func main() {
	fmt.Println("Time, let's discuss it")

	presentTime := time.Now() // Use the renamed package
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 59, 0, 0, time.UTC) // Use the renamed package
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
