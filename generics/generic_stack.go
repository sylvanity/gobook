package main

import "fmt"

// Stack is a generic LIFO data structure.
type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack.
// It returns the element and a boolean indicating if the operation was successful.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T // The zero value for the type T.
		return zero, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index] // Shrink the slice.
	return element, true
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	fmt.Println("--- Generic Stack[int] ---")
	// Instantiate a Stack for integers.
	var intStack Stack[int]
	fmt.Println("Is empty?", intStack.IsEmpty())
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Is empty?", intStack.IsEmpty())
	val, ok := intStack.Pop()
	fmt.Printf("Popped: %d, ok: %t\n", val, ok)
	val, ok = intStack.Pop()
	fmt.Printf("Popped: %d, ok: %t\n", val, ok)
	val, ok = intStack.Pop()
	fmt.Printf("Popped: %d, ok: %t (after empty)\n", val, ok)

	fmt.Println("\n--- Generic Stack[string] ---")
	// Instantiate a Stack for strings.
	var stringStack Stack[string]
	stringStack.Push("hello")
	stringStack.Push("world")
	str, ok := stringStack.Pop()
	fmt.Printf("Popped: %q, ok: %t\n", str, ok)
	str, ok = stringStack.Pop()
	fmt.Printf("Popped: %q, ok: %t\n", str, ok)
}
