package main

import "fmt"

func main() {

	num1 := 12
	num2 := 20

	var result bool

	//equal op
	result = (num1 == num2)

	fmt.Printf("%d == %d returns %t \n", num1, num2, result)

	//not equal op
	result = (num1 != num2)

	fmt.Printf("%d != %d returns %t \n", num1, num2, result)

	//greater than op
	result = (num1 > num2)

	fmt.Printf("%d > %d returns %t \n", num1, num2, result)

	//less than op
	result = (num1 < num2)

	fmt.Printf("%d < %d returns %t \n", num1, num2, result)
}
