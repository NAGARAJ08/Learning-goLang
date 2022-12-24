package main

import "fmt"

func main() {
	var sname string
	var zip int

	// take name and age input
	fmt.Println("Enter your School name and Zipcode:")
	fmt.Scanln(&sname, &zip)

	fmt.Printf("Name: %s\nAge: %d", sname, zip)

	//provide the input using space in between
}
