package main

import "fmt"

func main() {
	// Declares an array that can hold 3 integers. All elements are initialized to 0.
	var highScores [3]int
	fmt.Println("Initial high scores:", highScores)

	// Access and assign values to elements using a zero-based index.
	highScores[0] = 98
	highScores[1] = 95
	highScores[2] = 92
	fmt.Println("Updated high scores:", highScores)
	fmt.Println("First score:", highScores[0])

	// Using an array literal
	primesLiteral := [4]int{2, 3, 5, 7}
	fmt.Println("Primes (literal):", primesLiteral)

	// Using '...' to let the compiler infer the length
	primesInferred := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primes (inferred length):", primesInferred)
	fmt.Println("Length of inferred primes:", len(primesInferred))
}
