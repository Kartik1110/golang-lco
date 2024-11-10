package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go!")

	var ptr *int // pointer to an integer
	fmt.Println("Value of ptr: ", ptr)

	myNumber := 23
	var ptr1 = &myNumber // & means reference to the memory location of myNumber
	fmt.Println("Value of myNumber: ", ptr1, *ptr1)

	*ptr1 = *ptr1 + 10
	fmt.Println("Value of myNumber: ", ptr1, *ptr1)
}
