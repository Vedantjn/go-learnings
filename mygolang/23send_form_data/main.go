package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("We'll create a simple POST request")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "vedant")
	data.Add("lastname", "jain")
	data.Add("email", "vedant@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
