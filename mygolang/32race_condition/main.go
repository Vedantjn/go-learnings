package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	// Initialize WaitGroup and RWMutex
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	// Initialize score slice with a single element
	var score = []int{0}

	// Add 4 to the WaitGroup counter
	wg.Add(4)

	// Goroutine 1: Append 1 to score
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// Goroutine 2: Append 2 to score
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// Goroutine 3: Append 3 to score
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// Goroutine 4: Read and print the score
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Four R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final score
	fmt.Println(score)
}
