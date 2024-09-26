package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	println("Welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, error := ioutil.ReadFile(filename)
	checkNilError(error)

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
