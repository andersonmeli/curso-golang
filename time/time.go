package main

import (
	"fmt"
	"time"
)

func main() {
	dateString1 := "2020-12-01T11:07:00.000Z"
	time1, err := time.Parse(time.RFC3339,dateString1)
	if err!=nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println("Time1: ", time1)

	dateString2 := "2020-12-01T11:27:00.000Z"
	time2, err := time.Parse(time.RFC3339,dateString2)
	if err!=nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println("Time2: ", time2)

	// Using time.Before() method
	g1 := time2.Before(time1)
	fmt.Println("time2 before time1:", g1)

	// Using time.After() method
	g2 := time2.After(time1)
	fmt.Println("time2 after time1:", g2)
}
