package main

import (
	"fmt"
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
}
