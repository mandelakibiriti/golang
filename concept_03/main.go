package main

import (
	"fmt"
	"time"
)

func getConfig () (bool, string, time.Time) {
	return false, "info",  time.Now()
}

func main() {
	// Assign variables as types only
	var start, middle, end float32
	fmt.Println(start, middle, end)

	// Assign variables as types and values
	var name, age, location = "John", 35, "New York"
	fmt.Println(name, age, location)

	// Assign variables via a function
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)

}