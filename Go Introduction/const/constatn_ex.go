package main

import "fmt"

const PRODUCT string = "RBC"
const PRICE = 500

func main() {

	fmt.Println(PRODUCT)
	fmt.Println(PRICE)

	//type of the constants
	fmt.Printf("%T\n", PRODUCT)
	fmt.Printf("%T\n", PRICE)

}
