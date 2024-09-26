package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	// Welcome message
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing the URL
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// Displaying URL components
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// Handling query parameters
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	// Iterating over query parameters
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// Constructing a URL
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=vedant",
	}

	// Converting URL struct to string
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
