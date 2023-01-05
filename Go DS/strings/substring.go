package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Go lang  Programming"
	substring1 := "Go lang"
	substring2 := "language"

	res := strings.Contains(text, substring1)
	fmt.Println(res)

	res = strings.Contains(text, substring2)
	fmt.Println(res)
}
