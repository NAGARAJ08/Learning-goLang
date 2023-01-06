package main

import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

type Triangle struct {
	base, height float32
}

func (t Triangle) area() float32 {
	return t.base * t.height * 0.5
}

func calculate(s Shape) float32 {
	return s.area()
}

func main() {

	r := Rectangle{10, 4}
	t := Triangle{8, 12}

	rectanleArea := calculate(r)
	fmt.Println("Area of Rectangle: ", rectanleArea)

	triangleArea := calculate(t)
	fmt.Println("Area of Rectangle: ", triangleArea)
}
