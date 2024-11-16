package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang!")

	greeter("Kartik")

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(2, 4, 8, 9)
	fmt.Println("Pro result is: ", proResult)
}

func greeter(name string) {
	fmt.Println("Namaste from Golang!, " + name)
}

func adder(val1 int, val2 int) int { // return type
	return val1 + val2
}

func proAdder(values ...int) int { // accepting multiple values - variadic function
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
