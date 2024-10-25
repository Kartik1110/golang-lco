package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	createdDate := time.Date(1999, time.October, 11, 0, 0, 0, 0, time.UTC)

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format(("2006-01-02 15:04:05 Monday")))

	fmt.Println("Created Date: ", createdDate.Format(("2006-01-02 15:04:05 Monday")))
}
