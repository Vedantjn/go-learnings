package main

import "fmt"

// jwtToken := 300000 // not allowed outside the function
// var  jwtToken int = 300000 	// allowed outside the function

// first letter capital means public
const LoginToken string = "alsdnflasdnfm"

func main() {
	var username string = "Vedant"
	fmt.Println("Hello", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4566549797
	var smallFloat1 float64 = 255.4566549797
	fmt.Println(smallFloat)
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("LoginToken if of type %T \n", LoginToken)
}
