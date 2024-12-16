package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter by 1 when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)

	// Simulating some work with sleep
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a wait group

	numWorkers := 3
	wg.Add(numWorkers) // Set the counter to the number of workers

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg) // Start a worker goroutine
	}

	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers are done")
}
