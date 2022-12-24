package main

import "fmt"

func main() {

	var floatvalue float32 = 5.42

	//type conversion
	var intvalue int = int(floatvalue)

	fmt.Printf("Float value is %g\n", floatvalue)
	fmt.Printf("integer value is %d", intvalue)
}
