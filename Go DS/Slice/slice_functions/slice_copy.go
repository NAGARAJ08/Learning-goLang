package main

import "fmt"

func main() {

	primenumbers := []int{2, 3, 5, 7}

	numbers := []int{1, 2, 3}

	copy(numbers, primenumbers)

	fmt.Println("Numbers: ", numbers)
}
