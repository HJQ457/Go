package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	fmt.Println(a)

	var b Point
	//b = a //error
	b = a.(Point)
	fmt.Println(b)
}
