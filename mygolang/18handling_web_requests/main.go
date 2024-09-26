package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://google.com"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	// fmt.Println("Response is: ", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content is: ", content)
}
