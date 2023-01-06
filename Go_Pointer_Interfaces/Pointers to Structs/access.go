package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	person := Person{"WBC", 34}

	var ptr *Person

	ptr = &person

	fmt.Println("Name: ", ptr.name)
	fmt.Println("Age : ", ptr.age)
}
