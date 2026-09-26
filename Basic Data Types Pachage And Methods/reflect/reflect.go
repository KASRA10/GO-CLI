package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"GO", "PHP"}
	b := []string{"GO", "PHP"}

	// if a == b {
	// 	fmt.Println("a and b are equal")
	// } //! error: invalid operation: a == b (slice can only be compared to nil)

	if a == nil || b == nil {
		fmt.Println("a or b is nil")
	} // nothing will be printed

	fmt.Println(reflect.DeepEqual(a, b))   // true
	fmt.Println(reflect.DeepEqual(a, nil)) // false

	//* Map
	user1 := map[string]any{
		"name":   "Kasra",
		"skills": []string{"Go", "PHP"},
	}

	user2 := map[string]any{
		"name":   "Kasra",
		"skills": []string{"Go", "PHP"},
	}

	fmt.Println(reflect.DeepEqual(user1, user2)) // true

	//! Use it mostly in tests. Be careful: nil and an empty slice are different.
	var a []string = nil
	b := []string{}

	fmt.Println(reflect.DeepEqual(a, b)) // false

	//* reflect.Copy copies elements from one slice or array to another.
	//* It is the reflection version of Go’s built-in copy.
	source := []int{10, 20, 30}
	destination := make([]int, 2)

	count := reflect.Copy(
		reflect.ValueOf(destination),
		reflect.ValueOf(source),
	)

	fmt.Println(count)       // 2
	fmt.Println(destination) // [10 20]

	//? For normal code, use the built-in version instead:
	//* * count := copy(destination, source)

	//* Swapper returns a function that swaps two indexes in a slice.
	//* It is useful when writing generic sorting or reordering logic.
	skills := []string{"PHP", "Go", "JavaScript"}

	swap := reflect.Swapper(skills)

	swap(0, 2)

	fmt.Println(skills) // [JavaScript Go PHP]
}
