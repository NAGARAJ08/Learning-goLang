package main

import (
	"fmt"
	"strings"
)

func main() {

	text1 := "go is fun to learn"

	text1 = strings.ToUpper(text1)

	fmt.Println(text1)

	text2 := "GO LANG"
	text2 = strings.ToLower(text2)
	fmt.Println(text2)
}
