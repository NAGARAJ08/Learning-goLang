/*

Slice is a collection of similar types of data, just like arrays.
But here slice does not have a fixed size we can add more elements

Sysntx:
numbers := []int{1, 2, 3, 4, 5}

NOTE:
numbers := [5]int{1, 2, 3, 4, 5} this becomes array
*/

package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
}
