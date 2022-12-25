package main

import "fmt"

func main() {
	NoOfDays := 28

	switch {
	case 28 == NoOfDays:
		fmt.Println("Its February")
	default:
		fmt.Println("Not February")
	}
}
