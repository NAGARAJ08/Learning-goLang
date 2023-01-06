package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	Person1 := Person{"WBC", 30}

	//create a struct type pointer that stores the address of person1
	var ptr *Person
	ptr = &Person1

	fmt.Println(Person1)
	fmt.Println(ptr)
}
