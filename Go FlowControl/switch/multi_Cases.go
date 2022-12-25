package main

import "fmt"

func main() {
	week := "sunday"

	switch week {

	case "saturday", "sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("WeeDay")
	default:
		fmt.Println("Invalid Day")
	}
}
