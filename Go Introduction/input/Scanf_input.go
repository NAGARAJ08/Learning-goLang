package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Println("Enter your name and age: ")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Name : %s \n Age : %d", name, age)
}
