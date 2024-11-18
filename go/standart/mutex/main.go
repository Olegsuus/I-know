package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex // Mutex to protect shared data
	value int
}

// Increment method to safely increment the counter
func (c *Counter) Increment() {
	c.mu.Lock()   // Lock the mutex before accessing the shared resource
	c.value++     // Increment the counter
	c.mu.Unlock() // Unlock the mutex after access
}

// Get method to safely access the counter value
func (c *Counter) Get() int {
	c.mu.Lock()         // Lock the mutex before reading the shared resource
	defer c.mu.Unlock() // Ensure the mutex is unlocked at the end of the function
	return c.value      // Return the current value of the counter
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Number of goroutines
	numGoroutines := 10

	// Launch multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			for j := 0; j < 1000; j++ {
				counter.Increment() // Safely increment the counter
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final counter value:", counter.Get())
}
