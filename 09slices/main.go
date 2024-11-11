package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go!")

	var fruitsList = []string{"Apple", "Banana", "Cherry"}
	fruitsList = append(fruitsList, "Dates")

	fmt.Printf("Fruits list: %T\n", fruitsList)
	fmt.Println("Fruits list:", fruitsList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555) // append a new element to the slice

	sort.Ints(highScore) // sort the slice in ascending order

	fmt.Printf("High scores: %T\n", highScore)
	fmt.Println("High scores:", highScore)

	var courses = []string{"Python", "JavaScript", "Golang", "Rust", "Java"}
	fmt.Println("Courses:", courses)

	courses = append(courses[:2], courses[3])
	fmt.Println("Courses:", courses)
}
