package main

import "fmt"

type Employee struct {
	name     string
	age      int
	isRemote bool
}

func (e *Employee) updateName(newName string) {
	e.name = newName
}

func main() {
	fmt.Println("Structs in Golang!")

	employee1 := Employee{
		name:     "Kartik",
		age:      25,
		isRemote: false,
	}

	employee1.updateName("John")

	fmt.Println("Employee name", employee1.name)

	job := struct {
		title  string
		salary int
	}{
		title:  "Software Engineer",
		salary: 100000,
	}

	fmt.Println("Job", job.title, job.salary)

}
