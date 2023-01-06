package main

import "fmt"

type Weather struct {
	city string
	temp int
}

func main() {

	weather := Weather{"WAKANDA", 24}
	fmt.Println("Initial Weather: ", weather)

	ptr := &weather
	ptr.temp = 25

	fmt.Println("Updated weather: ", weather)
}
