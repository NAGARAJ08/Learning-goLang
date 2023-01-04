package main

import "fmt"

/*
syntax

var array_variable = [size]datatype{elements of array}

*/
func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//without size
	fmt.Println("")
	fmt.Println("without size")
	var Arr = [...]string{"Hello", "World", "!"}
	fmt.Println(Arr)

	//shorthand notation
	fmt.Println("")
	fmt.Println("shorthand notation")
	Array := [...]string{"Hello", "World", "!", "Shorthand"}
	fmt.Println(Array)

	//Accessing elements
	fmt.Println("")
	fmt.Println("Accessing elements")
	fmt.Println(arr[3])
	fmt.Println(Arr[1])
	fmt.Println(Array[3])

	//Array initialization
	fmt.Println("")
	fmt.Println("Array initialization")
	var arrayinit [4]int

	arrayinit[0] = 10
	arrayinit[1] = 20
	arrayinit[2] = 30
	arrayinit[3] = 40

	fmt.Println(arrayinit)

	//Initialize specific elements of an array
	fmt.Println("")
	fmt.Println("Initialize specific elements of an array")
	arrayinitSpecific := [5]int{0: 20, 3: 40}
	fmt.Println(arrayinitSpecific)
	fmt.Println("")

	//change the element
	fmt.Println("Before array")
	fmt.Println(arrayinit)
	arrayinit[2] = 300
	fmt.Println("After changing")
	fmt.Println(arrayinit)

	//length of the array
	length := len(arrayinitSpecific)
	fmt.Println("The Length of the array is", length)
	fmt.Print(arrayinitSpecific)
}
