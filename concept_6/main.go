package main

import (
	"fmt"
	"time"
)

func main()  {
	// Declare a pointer using a var statement and initialize it:
	count1 := new(int)
	// get some memory for a type and return a pointer to that address
	count2 := new(int)
	// Using &, create a pointer from the existing variable
	countTemp := 5
	count3 := &countTemp
	// create a pointer from some types without a temporary variable.
	t := &time.Time{}
	
	fmt.Printf("count1: %#v\n", count1)
	if count1 != nil {
		// Dereference the pointer to get the value
		fmt.Printf("Value of count1: %#v\n", *count1)
	} else {
		fmt.Println("count1 is nil")
	}

	fmt.Printf("count2: %#v\n", count2)
	// Dereference the pointer to get the value
	fmt.Printf("Value of count2: %#v\n", *count2) 

	fmt.Printf("count3: %#v\n", count3)
	// Dereference the pointer to get the value
	fmt.Printf("Value of count3: %#v\n", *count3)

	fmt.Printf("time : %#v\n", t)
	if t != nil {
		// Dereference the pointer to get the value
		fmt.Printf("Value of time : %#v\n", *t)
	} else {
		fmt.Println("time is nil")
	}
}