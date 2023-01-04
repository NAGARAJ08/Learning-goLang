package main

import "fmt"

func main() {

	//create a slice from an array

	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	sliceNum := numbers[4:7]

	fmt.Println(sliceNum)
}
