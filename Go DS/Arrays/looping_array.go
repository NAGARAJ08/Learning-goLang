package main

import "fmt"

func main() {

	age := [...]int{12, 13, 12, 10}

	for i := 0; i < len(age); i++ {
		fmt.Print(age[i])
		fmt.Print(" ")
	}

	//multidimensional arrays
	fmt.Println("")
	arrayIntegers := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(arrayIntegers[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
