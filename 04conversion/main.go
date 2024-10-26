package main

import (
	"bufio" // used to read user inputs
	"fmt"
	"os"      // used to interact with the OS
	"strconv" // used to convert data types
	"strings" // used to manipulate strings
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating our Pizza: ", input)

	// Conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Converted number: ", numRating)
		fmt.Printf("Variable is of type: %T", numRating)
	}
}
