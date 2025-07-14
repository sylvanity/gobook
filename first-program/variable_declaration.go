package main

import "fmt"

func main() {
	// Explicit declaration with type
	var score1 int = 100
	fmt.Println("score1:", score1)

	// Type inference
	var score2 = 100 // Type int is inferred by the compiler.
	fmt.Println("score2:", score2)

	// Short variable declaration (inside function)
	// This is equivalent to 'var score = 100' but can only be used inside functions.
	score3 := 100
	fmt.Println("score3:", score3)

	// Zero value
	var score4 int
	fmt.Println("score4 (zero value):", score4)
}
