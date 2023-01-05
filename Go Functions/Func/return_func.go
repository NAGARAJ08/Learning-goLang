package main

import "fmt"

func addNum(num1 int, num2 int) int {

	sum := num1 + num2
	return sum
}
func main() {

	res := addNum(21, 21)
	fmt.Println("Sum :", res)
}
