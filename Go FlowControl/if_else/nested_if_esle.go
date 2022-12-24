package main

import "fmt"

func main() {
	num1 := 12
	num2 := 20

	if num1 >= num2 {

		if num1 == num2 {
			fmt.Printf("Result: %d == %d", num1, num2)
		} else {
			fmt.Printf("Result: %d > %d", num1, num2)
		}
	} else {
		fmt.Printf("Result: %d < %d", num1, num2)
	}
}
