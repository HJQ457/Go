package main

import "fmt"

func main() {
	fmt.Println(10 / 4)
	//结果为2
	//说明：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	//如果希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// % 的使用
	fmt.Println("10 % 3 =", 10%3)

	// ++ 和 --
	var i int = 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}
