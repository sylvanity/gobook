package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Incorrect Usage (capturing loop variable) ---")
	// Incorrect usage! All goroutines capture the same *i* variable.
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second) // Wait for goroutines to likely finish

	fmt.Println("\n--- Correct Usage (passing loop variable as argument) ---")
	for i := 0; i < 5; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i) // Pass 'i' as an argument
	}
	time.Sleep(time.Second) // Wait for goroutines to likely finish
}
