package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	// Create a buffered channel with capacity 2
	myCh := make(chan int, 2)
	// Create a WaitGroup to synchronize goroutines
	wg := &sync.WaitGroup{}

	// Add 2 to the WaitGroup counter
	wg.Add(2)

	// Receiver goroutine (receive-only channel)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// Receive a value from the channel and check if it's open
		val, isChannelOpen := <-myCh

		// Print the channel status and received value
		fmt.Println("Is channel open?", isChannelOpen)
		fmt.Println("Received value:", val)

		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Sender goroutine (send-only channel)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// Send a value to the channel
		myCh <- 0
		// Close the channel after sending
		close(myCh)
		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Wait for both goroutines to finish
	wg.Wait()
}

// Definitions:
// Channel: A way to send and receive data between goroutines
// Buffered Channel: A channel that can hold multiple values
// WaitGroup: A synchronization primitive that allows you to wait for a collection of goroutines to finish
// Goroutine: A lightweight thread managed by the Go runtime
