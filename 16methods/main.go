package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang!")

	kartik := User{"Kartik", "kartik@gmail.com", 25, true}

	fmt.Println("Name: ", kartik.Name)
	kartik.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active? ", u.Status)
}
