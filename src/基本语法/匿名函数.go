package main

import "fmt"

func main() {

	//匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 = ", res1)

	//将匿名函数赋给a
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2 = ", res2)
}
