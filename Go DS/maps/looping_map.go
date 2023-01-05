package main

import "fmt"

func main() {

	squared := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	for Number, sqsquared := range squared {
		fmt.Printf("Square of %d is %d \n", Number, sqsquared)
	}
}
