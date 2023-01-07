package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Printf("Please enter your string for processing: ")
	fmt.Scan(&s)
	ls := strings.ToLower(s)
	if strings.HasPrefix(ls, "i") && strings.Contains(ls, "a") && strings.HasSuffix(ls, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
