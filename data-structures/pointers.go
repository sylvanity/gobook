package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
}

func main() {
	fmt.Println("--- Pointer Basics ---")
	name := "Alice"
	p := &name // p is a pointer to a string. Its value is the memory address of 'name'.

	fmt.Println("Value of name:", name)
	fmt.Println("Memory address of name:", p)
	fmt.Println("Value at the address p points to:", *p)

	// Modify the value through the pointer
	*p = "Bob"
	fmt.Println("Value of name after modification via pointer:", name)

	fmt.Println("\n--- Pointers and Structs ---")
	// The idiomatic way to create a pointer to a struct instance.
	empPtr := &Employee{
		ID:        103,
		FirstName: "Peter",
	}
	fmt.Printf("Struct pointer: %p, value: %+v\n", empPtr, *empPtr)

	// Using new() - less common for structs
	empPtr2 := new(Employee)
	fmt.Printf("Struct pointer (from new): %p, value: %+v\n", empPtr2, *empPtr2)
}
