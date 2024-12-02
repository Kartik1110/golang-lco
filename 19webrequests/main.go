package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests in golang")

	PerformGetRequest()
	PerformJsonPostRequest()
}

func PerformGetRequest() {
	url := "http://localhost:3000/get"

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

func PerformJsonPostRequest() {
	const url = "http://localhost:3000/post-json"

	// json payload
	body := strings.NewReader(`
		{
			"name": "Kartik",
			"age": 25
		}
	`)

	res, err := http.Post(url, "application/json", body)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}
