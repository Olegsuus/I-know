package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64 // Using int64 for atomic operations

	// Number of goroutines
	numGoroutines := 10

	// Launch multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // Atomically increment the counter
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final counter value:", atomic.LoadInt64(&counter)) // Atomically load the counter value
}
