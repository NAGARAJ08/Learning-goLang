package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Hello!"
	str2 := "Hello! there"
	str3 := "Hello!"

	fmt.Println(strings.Compare(str1, str2)) //-1 bcz str1 < str2
	fmt.Println(strings.Compare(str2, str3)) //1 bcz str2 > str3
	fmt.Println(strings.Compare(str1, str3)) //0 str1 == str3
}
