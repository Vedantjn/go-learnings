package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// There is no inheritance in golang; No super or parent

	vedant := User{"Vedant", "vedant@go.dev", true, 22}

	fmt.Printf("Vedant's details are: %v\n", vedant)
	fmt.Printf("Vedant's details are: %+v\n", vedant)
	fmt.Printf("Name is %v and email is %v\n", vedant.Name, vedant.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
