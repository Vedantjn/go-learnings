package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// Declare and initialize integer and float variables
	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5

	// This line would cause a compilation error due to mismatched types
	// fmt.Println("The sum is:", mynumberOne + mynumberTwo) // invalid operation: mynumberOne + mynumberTwo (mismatched types int and float64)

	// Convert float to int before addition
	fmt.Println("The sum is:", mynumberOne+int(mynumberTwo))

	// Random number generation

	// Method 1 -> deprecated
	// rand.Seed(time.Now().UnixNano()) // deprecated
	// fmt.Println(rand.Intn(5))

	// Method 2
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(r.Intn(5))

	// Method 3 -> using crypto/rand for secure random number generation
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
