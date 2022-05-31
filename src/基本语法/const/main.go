package main

import "fmt"

func main() {
	var num int
	num = 10

	//常量声明时必须赋值
	const tax int = 0
	//常量不可以修改
	//常量只能是bool，int，float，string类型

	fmt.Println(num, tax)

	const (
		a = iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
