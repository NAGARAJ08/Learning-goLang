package main

import "fmt"

func fact(num int) int {

	if num == 0 {
		return 1
	} else {
		return num * fact(num-1)
	}
}
func main() {
	num := 5
	res := fact(num)

	fmt.Println("The factorial of the number is: ", res)
}
