package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Tags     []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	course := Course{
		Name:     "100xDevs",
		Price:    5499,
		Platform: "youtube",
		Tags:     []string{"Web dev", "Web 3"},
	}

	jsonData, err := json.MarshalIndent(course, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}
