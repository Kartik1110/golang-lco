package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go!")

	// Declare an array of integers
	var numbers [5]int
	fmt.Println("Empty array:", numbers)

	// Initialize an array with values
	numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", numbers)

	// Access elements of the array
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])

	// Iterate over the array
	fmt.Println("Iterating over array:")
	for i, num := range numbers {
		fmt.Printf("Index %d: %d\n", i, num)
	}

	// Multi-dimensional array
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D array:", matrix)

	// Access elements of the 2D array
	fmt.Println("Element at [1][2]:", matrix[1][2])

	var fruits = [3]string{"Apple", "Banana", "Cherry"}

	fmt.Println("Fruits array:", fruits)
}
