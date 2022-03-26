package main

import "fmt"

type MethodUtils struct {
	long   int
	weight int
}

func (m MethodUtils) Print() {
	for i := 0; i < m.weight; i++ {
		for j := 0; j < m.long; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) Print2(w int, l int) {
	for i := 0; i < w; i++ {
		for j := 0; j < l; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var m MethodUtils
	m.long = 5
	m.weight = 3
	m.Print()

	fmt.Println()

	m.Print2(2, 3)
}
