package main

import "fmt"

//intitalize the func rectangle
type Rectangle func(int, int) int

//create a structure
type rectangleParam struct {
	length  int
	breadth int
	color   string

	//function as a field of struct
	rect Rectangle
}

func main() {

	result := rectangleParam{
		length:  10,
		breadth: 20,
		color:   "Black",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println("Color of Rectangle: ", result.color)
	fmt.Println("Area of Rectangle: ", result.rect(result.length, result.breadth))
}
