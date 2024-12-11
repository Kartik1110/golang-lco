package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}
