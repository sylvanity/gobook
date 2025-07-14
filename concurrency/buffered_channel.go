package main

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer sending: %d\n", i)
		ch <- i
	}
	fmt.Println("Producer closing channel.")
	close(ch) // Close the channel when done sending.
}

func main() {
	ch := make(chan int, 5)
	go producer(ch)

	fmt.Println("Consumer waiting for values...")
	// This loop receives values from the channel until it is closed and empty.
	for v := range ch {
		fmt.Printf("Consumer received: %d\n", v)
	}
	fmt.Println("Consumer detected channel closed. Exiting.")
}
