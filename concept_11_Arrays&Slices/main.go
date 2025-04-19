package main

import (
	"fmt"
	"os"
)

// Arrays
func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}
	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
	arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
	arr[i] = arr[i] * arr[i]
	}
	return arr
}

// Slices
func genSlices() ([]int, []int, []int, []int) {
	// slice1 is nil
	var slice1 []int
	// slice2 is an empty slice
	slice2 := []int{}
	// slice3 is a slice with 5 elements
	slice3 := []int{1, 2, 3, 4, 5}
	// slice4 with lenght 10 and capacity 50
	slice4 := make([]int, 10, 50)

	return slice1, slice2, slice3, slice4
}

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		return nil
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func getPassedLocalesArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}
func getLocales(extraLocales []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)
	return locales
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	// append of a new value above the capacity of the slice 
	// which will cause the slices s1 and s2 to be delinked
	// Hence the values of s1 and s2 will be different
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append(s1[:0:0], s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	/*
		Arrays
	*/
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0} :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)

	// the value is false because the last element is different as the default value
	// of the array is 0 for elements in the array
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3) 

	var array1 [10]int
	// set key 9 to value 0
	array2 := [...]int{9: 1}
	// set key 0 to value 1, set key 9 to value 10,
	// and set key 4 to value 5
	array3 := [10]int{1, 9: 10, 4: 5}
	fmt.Println(array1, array2, array3)

	// Reading from an array
	arrString := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	fmt.Println(arrString[1], arrString[0], arrString[3], arrString[2])

	// Writing to an array
	arrString[1] = "It’s"
	arrString[0] = "time"
	fmt.Println(arrString[1], arrString[0], arrString[3], arrString[2])

	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)

	/*
	 	Slices
	*/
	s1, s2, s3, s4 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
	fmt.Printf("s3: len = %v cap = %v\n", len(s4), cap(s4))

	// Creating a slice from an array
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First :", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last :", s[len(s)-1], s[len(s)-1:])
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4 :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	fmt.Println(m)

	// Controlling internal slice behavior
	l1, l2, l3 := linked()
	fmt.Println("Linked        :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link       :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link      :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link   :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link  : ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)

	// execute code in terminal: go run . Run go run whatever args you want to run
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
	// append values to slice
	locales := getLocales(getPassedLocalesArgs())
	fmt.Println("Locales to use:", locales)

}
