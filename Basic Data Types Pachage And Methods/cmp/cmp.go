package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	x := 10
	z := 100

	fmt.Println(cmp.Compare(x, z))              // -1
	fmt.Println(cmp.Compare("apple", "banana")) // -1

	scores := []int{80, 20, 100, 50}

	slices.SortFunc(scores, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	fmt.Println(scores) // [20 50 80 100]

	/*
		cmp.Less
		cmp.Less(a, b) returns true when a should come before b; it is similar to writing a < b.
	*/

	fmt.Println(cmp.Less(5, 10))       // true
	fmt.Println(cmp.Less(10, 5))       // false
	fmt.Println(cmp.Less("Go", "PHP")) // true

	/*
		cmp.Or
		cmp.Or returns the first value that is not the zero value.
		If all values are zero values, it returns the zero value. This is useful for fallback/default values
	*/
	name := cmp.Or("", "Kasra", "Guest")

	fmt.Println(name) // Kasra

	port := cmp.Or(0, 8080)

	fmt.Println(port) // 8080

	enabled := cmp.Or(false, true)

	fmt.Println(enabled) // true
}
