package main

import "fmt"

func main() {

	var temperature float32
	var sunny bool
	fmt.Println("Enter the Current Temperature:")
	fmt.Println("Is the Day Sunny")
	fmt.Scanf("%g %t", &temperature, &sunny)

	//fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature is :%g \n Is the day Sunny? %t", temperature, sunny)
}
