package main

import "fmt"

func sum(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + sum(num-1)
	}
}
func main() {

	num := 50
	res := sum(num)
	fmt.Println("Sum:", res)

}
