// Program to return the area of a rectangle

package main

import "fmt"

func main() {

	area := func(length, breadth int) int {
		return length * breadth
	}

	fmt.Println("The area of rectangle is", area(3, 4))

}
