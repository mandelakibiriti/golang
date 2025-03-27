package main

import (
	"fmt"
)

func main() {
	
	// Main Course
	var total float64 = 13 * 2
	fmt.Println("Total: ", total)
	
	// Drinks
	var totalDrinks float64 = 4 * 2.25
	fmt.Println("Total Drinks: ", totalDrinks)
	total = total + totalDrinks
	fmt.Println("Total: ", total)
	
	// total tip (10%)
	var tip float64 = total * 0.1
	fmt.Println("Tip: ", tip)
	tip += 5
	fmt.Println("New Tip: ", tip)
	
	// discount
	discount := tip
	discount -= 5
	fmt.Println("Discount: ", discount)
	
	// total bill
	total = total + tip
	fmt.Println("Total: ", total)
	
	// split bill
	split := total / 2
	fmt.Println("Split: ", split)
	
	// Reward every 5th visit
	visit := 24
	visit ++ 
	var remainder  = visit % 5
	if remainder == 0 {
		fmt.Println("Reward given for your visit")
	}
	fmt.Println("Current Visit: ", visit)
	oldVisit := visit
	oldVisit--
	fmt.Println("Old Visit: ", oldVisit)
	var newRemainder = oldVisit % 5
	if newRemainder != 0 {
		fmt.Println("No Reward given for your visit")
	}
}