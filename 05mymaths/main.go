package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Maths in Go")

	var numberOne int = 2
	var numberTwo float64 = 3

	fmt.Println("Addition: ", numberOne+int(numberTwo))

	//random number
	rand.Seed(time.Now().UnixNano()) // provide a seed value
	fmt.Println(rand.Intn(100))
}
