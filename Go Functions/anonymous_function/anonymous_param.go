package main

import "fmt"

func main() {

	sum := func(num1 int, num2 int) {
		sum := num1 + num2
		fmt.Println("Sum is:", sum)
	}

	sum(100, 400)
}
