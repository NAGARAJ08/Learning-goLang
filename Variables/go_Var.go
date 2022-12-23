package main

import "fmt"

func main() {

	//explicitly declare data type
	var num_1 int = 10
	fmt.Println(num_1)

	//assign a value and dont declare the type
	var num_2 = 20
	fmt.Println(num_2)

	//shorthadn notation to define var
	num_3 := 30
	fmt.Println(num_3)

	num := 100
	fmt.Println("Initial Assignment ", num)

	num = 1000
	fmt.Println("Changed value: ", num)

	//Creating multiple var at once
	var name, age = "ABC", 22
	fmt.Println("Name is :", name, "and Age is: ", age)

	//Constants in go
	const lightspeed = 299792458

	//lightspeed = 2997924502 cannot be changed throws error

}
