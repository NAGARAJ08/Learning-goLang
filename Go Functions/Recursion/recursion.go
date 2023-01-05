package main

import "fmt"

func countDown(num int) {

	if num > 0 {
		fmt.Println(num)
		countDown(num - 1)
	} else {
		fmt.Println("Count Stop")
	}
}
func main() {
	countDown(20)
}
