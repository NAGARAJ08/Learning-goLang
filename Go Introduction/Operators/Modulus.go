package main

import "fmt"

func main() {

	c := 11
	d := 4
	fmt.Println("Modulus Operator")
	remainder := c % d
	fmt.Printf("%d Modulus %d = %d ", c, d, remainder)
}
