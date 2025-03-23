package main

import (
	"fmt"
	"time"
)

func main() {
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

	/*
		when passing by value, the changes you make to the value in the
		add5Value() function do not affect the value of the count variable that’s passed
		to the main function. While passing a pointer to the add5Pointer() function
		a value does change the value of the count variable passed to the main function
	*/
	var count int
	add5Value(count)
	fmt.Println("Count initial value: ", count)

	add5Pointer(&count)
	fmt.Println("Count Value from Pointer: ", count)

	// Swap values using pointers
	a, b := 5, 10
	fmt.Println("Values before swap: ", a, b)
	// Take addresses of a and b and pass into functions as pointers
	valueSwap(&a,&b)
	fmt.Println("Values after swap: ", a, b)
}

func add5Value(count int) {
	count += 5
	fmt.Println("Count Value: ", count)
}

func add5Pointer(count *int) {
	*count += 5
	fmt.Println("Count Pointer value: ", *count)
}

func valueSwap(a *int, b *int)  {
	*a,*b = *b, *a
}