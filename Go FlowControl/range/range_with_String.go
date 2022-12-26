package main

import "fmt"

func main() {

	string := "Go lang"
	fmt.Println("Index: character")

	for i, item := range string {
		fmt.Printf("%d = %c \n", i, item)
	}
}
