package main

import "fmt"

func main() {
	week := 5

	switch week {

	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("wedday")
	case 5:
		fmt.Println("Thursday")
		fallthrough
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	}

}
