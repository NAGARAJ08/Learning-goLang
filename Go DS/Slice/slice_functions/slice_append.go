package main

import "fmt"

func main() {
	primenum := []int{2, 3}

	primenum = append(primenum, 5, 7)
	fmt.Println("Prime Numbers: ", primenum)

}
