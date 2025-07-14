package main

import (
	"fmt"
	"strconv"
)

// Map applies the function f to each element of slice s and returns a new slice of results.
func Map[T, V any](s []T, f func(T) V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements of s for which f returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	// Pre-allocating capacity can improve performance
	result := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("--- Generic Map Function ---")
	numbers := []int{1, 2, 3, 4}

	// Example 1: Squaring a slice of integers.
	squared := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared numbers:", squared)

	// Example 2: Converting a slice of integers to a slice of strings.
	asStrings := Map(numbers, func(n int) string {
		return "Number: " + strconv.Itoa(n)
	})
	fmt.Println("Numbers as strings:", asStrings)

	fmt.Println("\n--- Generic Filter Function ---")
	// Example: Filtering for even numbers.
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evens)
}
