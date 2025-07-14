package main

import "fmt"

func main() {
	fmt.Println("--- for...range with a slice ---")
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("Index %d has value %d\n", i, v)
	}

	fmt.Println("\n--- for...range with a slice (value only) ---")
	for _, v := range nums {
		fmt.Println("Value:", v)
	}

	fmt.Println("\n--- for...range with a map ---")
	userAges := map[string]int{"Alice": 30, "Bob": 25}
	for name, age := range userAges {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	fmt.Println("\n--- for...range with a string ---")
	s := "Go!"
	for i, r := range s {
		fmt.Printf("Byte index %d: rune %q\n", i, r)
	}

	fmt.Println("\n--- for...range with a channel ---")
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Received:", v)
	}
}
