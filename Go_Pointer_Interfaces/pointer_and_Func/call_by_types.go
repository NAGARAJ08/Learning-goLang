package main

import "fmt"

func callByValue(num int) {

	fmt.Println(num)
}

//call by Reference
func callByReference(num *int) {

	*num = 300
	fmt.Println(*num)
}
func main() {

	var number int = 100

	callByValue(number)

	callByReference(&number)

}
