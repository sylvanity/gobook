package main

import "fmt"

func main() {
	fmt.Println("--- Slicing an array ---")
	underlyingArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Underlying array:", underlyingArray)

	mySlice := underlyingArray[2:5] // Contains elements at indices 2, 3, and 4.
	fmt.Printf("Slice: %v, len: %d, cap: %d\n", mySlice, len(mySlice), cap(mySlice))

	fmt.Println("\n--- Modifying a slice element ---")
	mySlice[0] = 99
	fmt.Println("Slice after modification:", mySlice)
	fmt.Println("Underlying array after modification:", underlyingArray)

	fmt.Println("\n--- Creating a slice with make ---")
	scores := make([]int, 5, 10) // Type, length 5, capacity 10.
	fmt.Printf("Scores (initial): %v, len: %d, cap: %d\n", scores, len(scores), cap(scores))

	fmt.Println("\n--- Appending to a slice ---")
	scores = append(scores, 95, 100)
	fmt.Printf("Scores (after append): %v, len: %d, cap: %d\n", scores, len(scores), cap(scores))

	// Appending beyond capacity
	scores = append(scores, 1, 2, 3, 4)
	fmt.Printf("Scores (after append beyond capacity): %v, len: %d, cap: %d\n", scores, len(scores), cap(scores))
}
