package main

import "fmt"

func main() {

	subject_marks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks Obtained")

	for subject, marks := range subject_marks {
		fmt.Println(subject, ":", marks)
	}
}
