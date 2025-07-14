package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Non-generic MinInt
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Non-generic MinFloat64
func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// The pre-generics approach using interface{}
func MinInterface(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			if a < b {
				return a
			}
			return b
		}
	case float64:
		if b, ok := b.(float64); ok {
			if a < b {
				return a
			}
			return b
		}
	}
	// In a real-world scenario, you might return an error here.
	return nil
}

// The generic version.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("--- Non-Generic Functions ---")
	fmt.Println("MinInt(3, 5):", MinInt(3, 5))
	fmt.Println("MinFloat64(3.14, 2.71):", MinFloat64(3.14, 2.71))

	fmt.Println("\n--- Pre-Generics with interface{} ---")
	fmt.Println("MinInterface(3, 5):", MinInterface(3, 5))
	fmt.Println("MinInterface(3.14, 2.71):", MinInterface(3.14, 2.71))

	fmt.Println("\n--- Generic Function ---")
	fmt.Println("Min[int](3, 5):", Min[int](3, 5))
	fmt.Println("Min(3.14, 2.71):", Min(3.14, 2.71)) // Type inference
	fmt.Println("Min(\"apple\", \"banana\"):", Min("apple", "banana"))
}
