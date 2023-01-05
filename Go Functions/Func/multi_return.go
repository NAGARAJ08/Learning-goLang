package main

import "fmt"

func Calculate(num1 int, num2 int) (int, int) {

	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}
func main() {

	sum, diff := Calculate(21, 10)
	fmt.Println("Sum :", sum, "Difference: ", diff)
}
