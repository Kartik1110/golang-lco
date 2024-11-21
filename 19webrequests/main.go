package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev"

func main() {
	fmt.Println("Welcome to web requests in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
