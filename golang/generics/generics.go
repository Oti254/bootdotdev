package main

/**
* Generics allow us to use variables to refer to specific types
* Allowing abstraction reducing code duplication
* T is the name of the type parameter and it must match the any constraint
**/
func getLast[T any] (s []T) T {
	var myZero T
	if len(s) == 0 {
		return myZero
	}
	lastElem := len(s) - 1
	return s[lastElem]
}
