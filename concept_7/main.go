package main

import (
	"fmt"
)

func main() {
	/* 
		A type definition creates a new, distinct type with the same underlying type and 
		operations as the given type, and binds an identifier to it. The new type is called 
		a defined type. It is different from any other type, including the type it 
		is created from. In this case the Weekday type is a defined type of int.
	*/
	type Weekday int

	/* 
		iota is a predeclared identifier representing the untyped integer ordinal number 
		of the current const specification in a (usually parenthesized) const declaration. 
		It is zero-indexed.
		
		iota is reset to 0 whenever the reserved word const appears in the source code.
		iota is incremented by one after each ConstSpec. iota can be used to create a 
		set of related constants that increase in value as they	go down the list.
	*/ 
	
	const (
		Sunday Weekday = iota + 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	/* 
		The const keyword is used to create a constant value.
		The Weekday type is used to define the type of the constant.
		The iota keyword is used to create a set of related constants that 
		increase in value as they go down the list. The + 1 is used to increment the 
		value of the Weekday type by 1.
	*/
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}