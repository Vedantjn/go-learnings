package main

import "fmt"

func main() {

	// defer follows LIFO

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
