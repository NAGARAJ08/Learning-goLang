package main

import (
	"fmt"
)

func main() {

	fmt.Print("Hello, ")
	fmt.Print("World!!")

	//Print variables
	name := "RBC"
	fmt.Print("Name: ", name)

	//Println

	currsal := 65000
	fmt.Println(" HELLO ,", name)
	fmt.Println("Your Current Salary is: ", currsal)

	//Printf() for integer %d
	currage := 23
	fmt.Printf("Your age is : %d ", currage)

	//Printf() for float %g
	var sal = 65000.543
	fmt.Printf("Annual Salary: %g ", sal)

	fmt.Printf(" \n %s is %d Years Old.", name, currage)
}
