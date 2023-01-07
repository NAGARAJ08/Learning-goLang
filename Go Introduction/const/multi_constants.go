package main

import "fmt"

const (
	PRODUCT  = "MANGO MOBILE"
	QUANTUTY = 120
	PRICE    = 74.49
	STOCK    = true
)

func main() {

	fmt.Print(QUANTUTY)
	fmt.Printf(" \t\t%T\n", QUANTUTY)

	fmt.Print(PRICE)
	fmt.Printf(" \t\t%T\n", PRICE)

	fmt.Print(PRODUCT)
	fmt.Printf(" \t%T\n", PRODUCT)

	fmt.Print(STOCK)
	fmt.Printf(" \t\t%T\n", STOCK)

}
