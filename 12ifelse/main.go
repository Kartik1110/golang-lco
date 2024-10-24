package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang!")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Less than 10"
	} else {
		result = "Greater than 10"
	}

	fmt.Println("Result:", result)
}
