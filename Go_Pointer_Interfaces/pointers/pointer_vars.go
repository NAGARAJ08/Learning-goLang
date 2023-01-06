package main

import "fmt"

func main() {

	var name = "RBC"
	var ptr *string

	ptr = &name

	fmt.Println("Value of the Pointer is: ", ptr)
	fmt.Println("Address of the variable: ", &name)
	fmt.Println("The value pointed by ptr is: ", *ptr)
}
