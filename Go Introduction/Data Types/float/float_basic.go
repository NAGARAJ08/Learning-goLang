package main

import "fmt"

func main() {
	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)

	//size explicitly will be 64 bits
	salary := 5676.3

	fmt.Println(salary)
}

/*
output

50000.504
50000.503882901
5676.3

*/
