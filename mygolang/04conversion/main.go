package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out Pizza Shop")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
