/*
Syn:

type StructName struct{
	//structures definititons
}
*/

package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	// assign value to struct while creating an instance
	person1 := Person{"RBC", 34}
	fmt.Println(person1)

	//define an instance
	var person2 Person

	person2 = Person{
		name: "WBC",
		age:  31,
	}
	fmt.Println(person2)
}
