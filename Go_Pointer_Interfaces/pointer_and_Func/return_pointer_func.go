package main

import "fmt"

func main() {

	res := display()
	fmt.Println("Welcom to Learning", *res)
}

func display() *string {

	msg := "Go Lang"
	return &msg
}
