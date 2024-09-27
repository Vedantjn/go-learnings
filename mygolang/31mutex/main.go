package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// wg is a pointer to a WaitGroup used to synchronize goroutines
var wg sync.WaitGroup

var mut sync.Mutex

func main() {
	// List of websites to check
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://giksjnfksjfthub.com",
	}

	// Iterate through the website list and start a goroutine for each
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // Increment the WaitGroup counter
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns

	// Send GET request to the endpoint
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
