package main

import "fmt"

func main() {
	fmt.Println("Structs in Go!")

	kartik := User{"Kartik", "kartik@gmail.com", true, 25}
	fmt.Println(kartik)

	fmt.Printf("Struct User: %+v\n", kartik)

	broker := Broker{"Broker", "broker@gmail.com", true, 24, kartik}

	fmt.Println("Broker:", broker)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Broker struct {
	Name     string
	Email    string
	Listings bool
	Age      int
	User     User
}
