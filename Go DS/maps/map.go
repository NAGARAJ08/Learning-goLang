/*
Syn:
subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
*/
package main

import "fmt"

func main() {

	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}

	fmt.Println(subjectMarks)

	//Accessing the valuesof map
	fmt.Println("")
	fmt.Println("Accessing the values of map")
	fmt.Println(subjectMarks["Golang"])

	//changing the value of map
	fmt.Println("")
	fmt.Println("changing the values of map")
	subjectMarks["Golang"] = 92

	fmt.Println("Updated", subjectMarks)
}
