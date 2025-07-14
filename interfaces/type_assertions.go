package main

import "fmt"

type Book struct {
	BookTitle string
}

func (b *Book) Title() string {
	return b.BookTitle
}

func checkType(i interface{}) {
	fmt.Printf("--- Checking type for value: %v ---\n", i)
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an integer with value: %d\n", v)
	case string:
		fmt.Printf("It's a string with value: %q\n", v)
	case bool:
		fmt.Printf("It's a boolean with value: %t\n", v)
	case *Book:
		fmt.Printf("It's a pointer to a Book with title: %s\n", v.Title())
	default:
		fmt.Printf("It's an unknown type: %T with value: %v\n", v, v)
	}
}

func main() {
	fmt.Println("--- Type Assertions ---")
	var i interface{} = "hello"

	// Risky assertion - will panic if type is wrong
	s := i.(string)
	fmt.Println("Risky assertion successful:", s)

	// Safe 'comma, ok' assertion
	s, ok := i.(string)
	if ok {
		fmt.Println("'comma, ok' assertion successful:", s)
	} else {
		fmt.Println("'comma, ok' assertion failed.")
	}

	var j interface{} = 123
	_, ok = j.(string)
	if !ok {
		fmt.Println("'comma, ok' assertion failed as expected. 'j' is not a string.")
	}

	fmt.Println("\n--- Type Switches ---")
	checkType(42)
	checkType("a test string")
	checkType(true)
	checkType(&Book{BookTitle: "The Go Programming Language"})
	checkType(12.34)
}
