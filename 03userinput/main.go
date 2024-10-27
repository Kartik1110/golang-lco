package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go programming!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok idiom or err idiom
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)
	fmt.Printf("Variable is of type: %T", rating)
}
