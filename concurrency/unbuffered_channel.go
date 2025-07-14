package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("Worker goroutine starting work...")
		time.Sleep(2 * time.Second)
		fmt.Println("Worker finished work, sending message.")
		messages <- "ping" // Send a value into the channel. This will block until main receives.
		fmt.Println("Worker message sent.")
	}()

	fmt.Println("Main goroutine waiting for a message...")
	msg := <-messages // Block until a value is received from the channel.
	fmt.Println("Main goroutine received:", msg)
}
