package main

import "fmt"

func main() {
	//var i int
	//i = 10

	a := 9
	b := 2
	fmt.Printf("交换前，a = %v,b = %v\n", a, b)

	t := a
	a = b
	b = t
	fmt.Printf("交换后，a = %v,b = %v\n", a, b)

	//不使用临时变量
	a = a + b
	b = a - b //a + b - b => b = a
	a = a - b //a + b - a => a = b
	fmt.Printf("交换后，a = %v,b = %v\n", a, b)

	a += 17
	fmt.Printf("a = %v\n", a)

	var d int
	d = a
	d = 8 + 2*8
	d = test1() + 90
	fmt.Printf("d = %v\n", d)
}

func test1() int {
	return 90
}
