package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the function returns

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		// Pass a pointer to the WaitGroup
		go worker(i, &wg)
	}

	fmt.Println("Main goroutine waiting for workers to finish...")
	wg.Wait() // Block until the counter goes to zero

	fmt.Println("All workers completed. Main goroutine exiting.")
}
