package main

import (
	"fmt"
	"strings"
)

func main() {

	var message = "I am Learning Golang"

	splitthemsg := strings.Split(message, " ")
	fmt.Println(splitthemsg)
}
