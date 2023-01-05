package main

import "fmt"

func main() {

	language := "Go Lang"

	fmt.Printf("%c ", language[0]) //G

	fmt.Printf("%c ", language[3]) //G

	fmt.Printf("%c ", language[5]) //G

	//length of the string
	fmt.Println()
	length := len(language)
	fmt.Println("Length of string is: ", length)
}
