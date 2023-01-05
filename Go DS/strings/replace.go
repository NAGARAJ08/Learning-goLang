package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "CAR"
	fmt.Println("Old String: ", text)

	//replace it t
	replacetext := strings.Replace(text, "R", "T", 1)

	fmt.Println("New String: ", replacetext)

}
