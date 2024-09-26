package main

import "fmt"

func main() {
	println("Hello to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proResult, myMessage := proAdder(2, 5, 6, 7, 8, 9)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("Pro message is : ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Vedant here"
}

func greeter() {
	println("Namaste from golang")
}
