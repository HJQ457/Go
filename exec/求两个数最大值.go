package main

import "fmt"

func main() {
	var n1 int = 10
	var n2 int = 20
	var max int

	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("两个数最大值为：", max)

	var n3 = 45
	if n3 > max {
		max = n3
	}
	fmt.Println("三个数最大值为：", max)
}
