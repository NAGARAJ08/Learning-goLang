package main

import "fmt"

func update(num *int) {

	*num = 30
}
func main() {
	num := 55

	update(&num)

	fmt.Println("The number is: ", num)
}
