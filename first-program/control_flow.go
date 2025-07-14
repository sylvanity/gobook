package main

import "fmt"

func main() {
	fmt.Println("--- if-else ---")
	score := 75
	if score > 90 {
		fmt.Println("Excellent!")
	} else if score > 60 {
		fmt.Println("Good job.")
	} else {
		fmt.Println("Please see me after class.")
	}

	fmt.Println("\n--- if with short initializer ---")
	someSlice := make([]int, 12)
	if n := len(someSlice); n > 10 {
		fmt.Println("The slice is too long.")
	}

	fmt.Println("\n--- for loop ---")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n--- switch ---")
	day := "Sunday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Monday":
		fmt.Println("Time to work.")
	default:
		fmt.Println("It's a weekday.")
	}
}
