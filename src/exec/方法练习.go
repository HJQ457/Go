package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64 {
	return 3.14 * c.radius * c.radius
	//return 3.14 * (*c).radius * (*c).radius	go底层做了处理
}

func main() {
	var a Circle
	a.radius = 1
	fmt.Println(a.area())

	(&a).radius = 2
	fmt.Println(a.area2())
}
