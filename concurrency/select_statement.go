package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We need to receive two messages, one from each channel.
	for i := 0; i < 2; i++ {
		fmt.Println("Waiting for a message...")
		select {
		case msg1 := <-c1:
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from channel 2:", msg2)
		}
	}
}
