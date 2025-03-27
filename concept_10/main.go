package main

import (
	"fmt"
	"unicode"
	"runtime"
)

/*
	Number Types
*/
func integerMemory() {
	// var list []int // using int consumes > 400MB of Heap memory
	var list []int8 // using int8 instead of int saves memory
	for i := 0; i < 10000000; i++ {
		list = append(list, int8(i))
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Length of array: %d TotalAlloc (Heap) = %v MiB\n", len(list), m.TotalAlloc/1024/1024)
}

/*
	Bool Type (True or False)
*/
func passwordChecker(pw string) bool {
	// Make it UTF-8 safe
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	if len(pwR) > 15 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// Bool types
	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	// Number types
	integerMemory()

	// Float types
	var a int = 100
	var b float32 = 100
	var c float64 = 100
	fmt.Println((a / 3))
	fmt.Println((b / 3))
	fmt.Println((c / 3))

	// Wrap around 
	var x int8 = 125
	var y uint8 = 253
	for i := 0; i < 5; i++ {
		x++
		y++
		fmt.Println(i, ")", "int8", x, "uint8", y)
	}

	// Text Types
	comment1 := `This is the BEST
	thing ever!` // raw text
	comment2 := `This is the BEST\nthing ever!` // raw text
	comment3 := "This is the BEST\nthing ever!" // interpreted text
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	// Rune Types
	username := "Sir_King_Über"
	runes := []rune(username)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}