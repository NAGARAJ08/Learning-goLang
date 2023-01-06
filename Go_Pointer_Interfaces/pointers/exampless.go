package main

import "fmt"

func main() {

	var num int
	var ptr *int
	num = 22
	fmt.Println("Address of the num is: ", &num)
	fmt.Println("Value of the num: ", num)

	ptr = &num
	fmt.Println("\nAddress of the pointer ptr: ", ptr)
	fmt.Println("\nvalue of the pointer ptr: ", *ptr)

	num = 10
	fmt.Println("\nAddress of the pointer ptr: ", ptr)
	fmt.Println("\nvalue of the pointer ptr: ", *ptr)

	*ptr = 1
	fmt.Println("Address of the num is: ", &num)
	fmt.Println("Value of the num: ", num)
}
