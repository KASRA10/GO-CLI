package main

import (
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "Kasra", Age: 24},
		{Name: "Ali", Age: 17},
	}

	isAdult := func(user User) bool {
		return user.Age >= 18
	}

	numbers := []int{10, 20, 30, 40, 20}

	fmt.Println(slices.Contains(numbers, 25)) // false

	// First position, or -1 if missing
	fmt.Println(slices.Index(numbers, 20)) // 1
	fmt.Println(slices.Index(numbers, 25)) // -1

	// Last position, or -1 if missing
	// fmt.Println(slices.LastIndex(numbers, 20))

	// Are two slices exactly equal?
	fmt.Println(slices.Equal(numbers, []int{10, 20, 30, 40, 20})) // true

	// Compare lexicographically: -1, 0, or 1
	fmt.Println(slices.Compare([]int{1, 2}, []int{1, 3})) // -1

	fmt.Println(slices.ContainsFunc(users, isAdult)) // true
	fmt.Println(slices.IndexFunc(users, isAdult))    // 0

	languages := []string{"Go", "PHP", "JavaScript"}

	// Insert at index 1
	languages = slices.Insert(languages, 1, "Python")
	fmt.Println(languages) // [Go Python PHP JavaScript]

	// Delete indexes from 1 up to, but not including, 3
	languages = slices.Delete(languages, 1, 3)
	fmt.Println(languages) // [Go JavaScript]

	// Replace indexes 1 through 2 with two new values
	languages = slices.Replace(languages, 1, 2, "TypeScript", "Rust")
	fmt.Println(languages) // [Go TypeScript Rust]

	original := []string{"Go", "PHP"}

	// Clone: make an independent copy
	copied := slices.Clone(original)

	copied[0] = "Rust"

	fmt.Println(original) // [Go PHP]
	fmt.Println(copied)   // [Rust PHP]
	/*
	   Clone is important because assigning one slice to another normally copies only the slice header;
	   both variables may still share the same backing array.
	*/

	// Concat: combine many slices
	frontend := []string{"HTML", "CSS", "JavaScript"}
	backend := []string{"Go", "PHP"}

	stack := slices.Concat(frontend, backend)

	fmt.Println(stack) // [HTML CSS JavaScript Go PHP]

	// Repeat the entire slice
	fmt.Println(slices.Repeat([]string{"Go", "PHP"}, 2))
	// [Go PHP Go PHP]

	scores := []int{80, 25, 100, 60}

	slices.Sort(scores)
	fmt.Println(scores) // [25 60 80 100]

	slices.Reverse(scores)
	fmt.Println(scores) // [100 80 60 25]

	fmt.Println(slices.Min(scores)) // 25
	fmt.Println(slices.Max(scores)) // 100

	fmt.Println(slices.IsSorted(scores)) // false

	// Binary search
	// BinarySearch is fast, but the slice must already be sorted in ascending order.
	// It returns the index and whether the target was found.

	numbers := []int{10, 20, 30, 40, 50}

	index, found := slices.BinarySearch(numbers, 30)

	fmt.Println(index, found) // 2 true

	numbers := []int{10, 20, 30, 40, 50}

	index, found := slices.BinarySearch(numbers, 35)

	fmt.Println(index, found) // 3 false

	//* Compact removes adjacent duplicate values, so sort first if duplicates are scattered.
	tags := []string{"go", "go", "php", "php", "javascript"}

	tags = slices.Compact(tags)

	fmt.Println(tags) // [go php javascript]

	tags := []string{"php", "go", "php", "javascript", "go"}

	slices.Sort(tags)
	tags = slices.Compact(tags)

	fmt.Println(tags) // [go javascript php]

	// Clip reduces a slice’s capacity to its current length
	// Grow reserves capacity before many appends, which can reduce allocations.

	items := make([]string, 0, 2)

	items = slices.Grow(items, 10)
	items = append(items, "Go", "PHP")

	items = slices.Clip(items)

	fmt.Println(len(items), cap(items)) // 2 2

	/*
		! Go 1.23 iterator helpers
		! These work with range and are useful when you want to process a slice in a clean, iterator-based style.
	*/

	languages := []string{"Go", "PHP", "JavaScript"}

	// All: indexes and values in normal order
	for index, language := range slices.All(languages) {
		fmt.Println(index, language)
	}

	languages := []string{"Go", "PHP", "JavaScript"}

	// Backward: indexes and values in reverse order
	for index, language := range slices.Backward(languages) {
		fmt.Println(index, language)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	// Chunk into groups of 3
	for group := range slices.Chunk(numbers, 3) {
		fmt.Println(group)
	}

	// [1 2 3]
	// [4 5 6]
	// [7]

	numbers := []int{1, 2}

	numbers = slices.AppendSeq(numbers, slices.Values([]int{3, 4, 5}))

	fmt.Println(numbers) // [1 2 3 4 5]
	//? AppendSeq appends values from an iterator sequence:
}
