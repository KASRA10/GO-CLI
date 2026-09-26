package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	// --------------------------------------------------
	// maps.Clone
	// --------------------------------------------------

	settings := map[string]string{
		"theme": "dark",
		"lang":  "en",
	}

	// Clone creates an independent map copy.
	copiedSett := maps.Clone(settings)

	fmt.Println(copiedSett)
	// map[lang:en theme:dark]

	copiedSett["theme"] = "light"

	fmt.Println(settings)
	// map[lang:en theme:dark]

	fmt.Println(copiedSett)
	// map[lang:en theme:light]

	// --------------------------------------------------
	// maps.Copy
	// --------------------------------------------------

	langs := map[string]string{
		"API":    "GO",
		"Server": "PHP",
		"Script": "JS",
	}

	priority2 := map[string]string{
		"API":    ".NET",
		"Script": "PHP",
	}

	fmt.Println(priority2)
	// map[API:.NET Script:PHP]

	maps.Copy(priority2, langs)

	fmt.Println(priority2)
	// map[API:GO Script:JS Server:PHP]

	// --------------------------------------------------
	// maps.Equal
	// --------------------------------------------------

	a := map[string]int{
		"go":  1,
		"php": 2,
	}

	b := map[string]int{
		"go":  1,
		"php": 2,
	}

	c := map[string]int{
		"go":  1,
		"php": 3,
	}

	fmt.Println(maps.Equal(a, b))
	// true

	fmt.Println(maps.Equal(a, c))
	// false

	// --------------------------------------------------
	// maps.DeleteFunc
	// --------------------------------------------------

	scores := map[string]int{
		"Kasra": 95,
		"David": 90,
		"John":  75,
	}

	maps.DeleteFunc(scores, func(name string, score int) bool {
		return score < 80
	})

	fmt.Println(scores)
	// map[Kasra:95 David:90]

	// Useful for removing expired cache entries,
	// invalid form data, inactive users, or empty values.

	// --------------------------------------------------
	// maps.EqualFunc
	// --------------------------------------------------

	roleA := map[string]string{
		"role": "DEVELOPER",
	}

	roleB := map[string]string{
		"role": "developer",
	}

	same := maps.EqualFunc(roleA, roleB, func(x, y string) bool {
		return strings.EqualFold(x, y)
	})

	fmt.Println(same)
	// true

	// --------------------------------------------------
	// maps.All
	// --------------------------------------------------

	user := map[string]string{
		"name": "Kasra",
		"role": "Developer",
		"city": "Baku",
	}

	for key, value := range maps.All(user) {
		fmt.Println(key, "=", value)
	}

	// --------------------------------------------------
	// maps.Keys
	// --------------------------------------------------

	for key := range maps.Keys(user) {
		fmt.Println(key)
	}

	// --------------------------------------------------
	// maps.Values
	// --------------------------------------------------

	for value := range maps.Values(user) {
		fmt.Println(value)
	}

	// --------------------------------------------------
	// maps.Keys + slices.Collect
	// --------------------------------------------------

	scores = map[string]int{
		"Kasra": 95,
		"Ali":   80,
		"Sara":  90,
	}

	keys := slices.Collect(maps.Keys(scores))
	values := slices.Collect(maps.Values(scores))

	fmt.Println(keys)
	fmt.Println(values)

	// --------------------------------------------------
	// slices.Sort
	// --------------------------------------------------

	slices.Sort(keys)

	fmt.Println(keys)
	// [Ali Kasra Sara]

	// --------------------------------------------------
	// maps.Insert
	// --------------------------------------------------

	profile := map[string]string{
		"name": "Kasra",
	}

	extra := map[string]string{
		"city": "Baku",
		"role": "Developer",
	}

	maps.Insert(profile, maps.All(extra))

	fmt.Println(profile)
	// map[city:Baku name:Kasra role:Developer]

	// --------------------------------------------------
	// maps.Collect
	// --------------------------------------------------

	source := map[string]int{
		"go":  1,
		"php": 2,
	}

	newMap := maps.Collect(maps.All(source))

	fmt.Println(newMap)
	// map[go:1 php:2]
}
