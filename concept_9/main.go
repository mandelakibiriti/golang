package main

import (
	"fmt"
	"errors"
	"time"
	"math/rand"
	"os"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 != 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}

func main() {
	/* 
		if and else if statements only execute the block that matches the condition
		if the condition is true, the block is executed and the rest of the blocks are skipped
		if the condition is false, the block is skipped and the next block is checked
		if no block matches the condition, the else block is executed
	 	if there is no else block, the program continues to the next statement
	*/
	input := 20
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}

	/* 
		The switch statement is used to compare a value against multiple conditions 
		which is followed by a value to compare.
		The case statement is followed by a colon and the block of code to execute 
		if the case matches the switch value. The case statement can have multiple values 
		separated by commas. The default statement is executed if no case matches the switch 
		value. The default statement can be omitted.
		The fallthrough statement is used to execute the next case block
	*/

	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Error, day born not valid")
	}

	/* 
		The for loop is used to execute a block of code multiple times
		The for loop has three parts: initialization, condition, and post
		The initialization part is executed before the loop starts
		The condition part is checked before each iteration
		The post part is executed after each iteration
		The post part can be omitted, but the initialization and condition parts are required
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"Jim", "Jane", "Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	topWord := ""
	topCount := 0
	for key, value := range words {
		if value > topCount {
			topCount = value
			topWord = key
		}
	}
	fmt.Println("Most popular word:", topWord)
	fmt.Println("With a count of  :", topCount)

	// Fizz Buzz challenge
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Bubble Sort
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", nums)
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}
	fmt.Println("After :", nums)

	// Fibonacci sequence
	a, b := 0, 1
	fib := []int{a, b}

	for i := 2; i < 10; i++ {
		a, b = b, a+b
		fib = append(fib, b)
	}
	
	fmt.Println(fib)

	/*
		Break and Continue statements are used to control the flow of the loop
		The break statement is used to exit the loop
		The continue statement is used to skip the rest of the loop and start the next iteration
	*/

	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		fmt.Println(r)
	}

	/*
		The goto statement is used to jump to a label in the code
		The label is defined by an identifier followed by a colon
		The goto statement is followed by the label to jump to
		The goto statement should be used sparingly as it can make the code harder to read
	*/
	for {
		t := rand.Intn(8)
		if t%3 == 0 {
			fmt.Println("Skip - Not divisible by 3")
			continue
		} else if t%2 == 0 {
			fmt.Println("Stop - Even Number")
			goto STOP
		}
		fmt.Println(t)
	}
STOP:
	fmt.Println("Goto label reached")

	/*
		The defer statement is used to delay the execution of a function until the surrounding function returns
		The deferred function is executed in LIFO order
		The deferred function is executed even if the surrounding function panics
	*/
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Attempting to open file")
		fmt.Println("Error opening file:", err)
		defer fmt.Println("Closing file...")
		return
	}
	defer file.Close()	
	fmt.Println("File closed successfully") // This line will not be executed

	/*
		The panic statement is used to stop the normal execution of the program
		The panic statement is followed by an error message
		The panic statement can be used to handle unexpected errors
		The recover statement is used to catch a panic and resume normal execution
		The recover statement is used in a deferred function
	*/	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Panic!") // This line will cause the program to panic
	fmt.Println("This line will not be executed")
}