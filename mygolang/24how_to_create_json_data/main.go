package main

import (
	"encoding/json"
	"fmt"
)

// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as JSON data

	finalJson1, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(finalJson))
	fmt.Printf("%s\n", finalJson1)
	//MarshalIndent(lcoCourses, "", "\t")

	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson2)

}
