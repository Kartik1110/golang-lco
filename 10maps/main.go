package main

import "fmt"

func main() {
	fmt.Println("Maps in Go!")

	languages := make(map[string]string)

	languages["Python"] = "Guido van Rossum"
	languages["Ruby"] = "Yukihiro Matsumoto"
	languages["JavaScript"] = "Brendan Eich"
	languages["Go"] = "Robert Griesemer"

	delete(languages, "Ruby")

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
