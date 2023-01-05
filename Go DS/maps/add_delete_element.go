package main

import "fmt"

func main() {

	students := map[int]string{1: "RBC", 2: "WBC"}
	fmt.Println("Init Map : ", students)
	fmt.Println("")

	//add element with key 3
	students[3] = "Plasma"

	//add element with key 5
	students[4] = "JCB"
	fmt.Println("Updated Map: ", students)
	fmt.Println("")

	delete(students, 2)
	fmt.Println("Updated map after deleting:")
	fmt.Println(students)

}
