package main

import "fmt"

func main() {
	fmt.Println("--- Creating and using maps ---")
	// Using make
	userAges := make(map[string]int)

	// Using a literal
	userAges = map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Initial map:", userAges)

	// Add, update, and retrieve
	userAges["Charlie"] = 35 // Add a new key-value pair
	userAges["Alice"] = 31   // Update an existing value
	age := userAges["Bob"]   // Retrieve a value
	fmt.Println("Updated map:", userAges)
	fmt.Println("Bob's age:", age)

	fmt.Println("\n--- The 'comma, ok' idiom ---")
	age, ok := userAges["David"]
	if !ok {
		fmt.Println("David's age is not in the map.")
	} else {
		fmt.Println("David's age is", age)
	}

	age, ok = userAges["Alice"]
	if !ok {
		fmt.Println("Alice's age is not in the map.")
	} else {
		fmt.Println("Alice's age is", age)
	}

	fmt.Println("\n--- Deleting from a map ---")
	delete(userAges, "Bob")
	fmt.Println("Map after deleting Bob:", userAges)

	fmt.Println("\n--- Iterating over a map (order is not guaranteed) ---")
	for name, age := range userAges {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
