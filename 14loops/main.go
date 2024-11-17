package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang!")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto runThisCode
		}

		if rougeValue == 5 {
			break
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

runThisCode:
	fmt.Println("This is running a label!")
}
