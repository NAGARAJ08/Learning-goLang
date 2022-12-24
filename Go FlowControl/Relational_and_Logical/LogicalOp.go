package main

import "fmt"

func main() {

	num1 := 6
	num2 := 12
	num3 := 6

	var result bool

	result = (num1 > num2) && (num1 == num3)
	fmt.Printf("Result of AND operator is %t \n", result)

	result = (num1 > num2) || (num1 == num3)
	fmt.Printf("Result of OR operator is %t \n", result)

	result = !(num1 == num3)
	fmt.Printf("Result of AND operator is %t \n", result)

}
