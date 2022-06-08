package defaultmap_test

import (
	"fmt"

	"github.com/spenserblack/go-defaultmap"
)

func ExampleMap() {
	m := defaultmap.NewMap[string](func() string { return "I'm the default!" })
	m.Insert("exists", "Hello, World!")

	fmt.Println(m.Get("exists"))
	fmt.Println(m.Get("doesn't exist"))
	fmt.Println(m.GetOr("also doesn't exist", "I'm a one-time default!"))
	// Output:
	// Hello, World!
	// I'm the default!
	// I'm a one-time default!
}

func ExampleMap_Get() {
	m := defaultmap.NewMap[string](func() string { return "I'm the default!" })
	m.Insert("exists", "Hello, World!")

	fmt.Println(m.Get("exists"))
	fmt.Println(m.Contains("doesn't exist"))
	fmt.Println(m.Get("doesn't exist"))
	fmt.Println(m.Contains("doesn't exist"))
	// Output:
	// Hello, World!
	// false
	// I'm the default!
	// true
}

func ExampleMap_GetOr() {
	m := defaultmap.NewMap[string](func() string { return "I'm the default!" })
	m.Insert("exists", "Hello, World!")

	fmt.Println(m.GetOr("exists", "other default"))
	fmt.Println(m.Contains("doesn't exist"))
	fmt.Println(m.GetOr("doesn't exist", "other default"))
	fmt.Println(m.Contains("doesn't exist"))
	// Output:
	// Hello, World!
	// false
	// other default
	// false
}

func ExampleMap_Contains() {
	m := defaultmap.NewMap[string](func() string { return "I'm the default!" })
	m.Insert("exists", "Hello, World!")

	fmt.Println(m.Contains("exists"))
	fmt.Println(m.Contains("doesn't exist"))
	// Output:
	// true
	// false
}
