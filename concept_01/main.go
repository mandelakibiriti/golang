package main

import (
	"fmt"
	"log"
	"conversion/conv"

	human "github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("Hello World!")
	
	// Module Import
	var number uint64 = 12345679
	fmt.Println("Size file is:", human.Bytes(number))

	fmt.Println("String to Integer Conversion Example")

	// Example 1: Valid conversion
	num, err := conv.Conv("123")
	if err != nil {
		log.Fatalln("Error:", err) // This will stdout on the terminal
	} else {
		fmt.Println("Converted number:", num)
	}

	// Example 2: Invalid conversion
	num, err = conv.Conv("abc")
	if err != nil {
		log.Fatalln("Error:", err) // This will stdout on the terminal
	} else {
		fmt.Println("Converted number:", num)
	}
}