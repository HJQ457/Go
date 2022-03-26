package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rigntDown Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(r1)
}
