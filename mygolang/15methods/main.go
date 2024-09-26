package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// There is no inheritance in golang; No super or parent

	vedant := User{"Vedant", "vedant@go.dev", true, 22}

	fmt.Printf("Name is %v and email is %v\n", vedant.Name, vedant.Email)

	vedant.GetStatus()
	vedant.NewMail()

	fmt.Printf("Name is %v and email is %v\n", vedant.Name, vedant.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
