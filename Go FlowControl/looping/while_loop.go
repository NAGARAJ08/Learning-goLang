package main

/*
Syntax:

for condition{
	//code block
}

*/
import "fmt"

func main() {
	multi := 1

	for multi <= 10 {

		product := 5 * multi
		fmt.Printf("5 * %d = %d\n", multi, product)
		multi++
	}
}
