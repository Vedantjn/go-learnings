package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("We'll create a simple GET request")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(content);
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}
