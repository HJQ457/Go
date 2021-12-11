package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//当直接输出byte值，就输出了对应字符的码值
	fmt.Println("c1=", c1, "c2=", c2)

	//可以采取格式化输出方式
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	var c3 = '北'
	fmt.Printf("c3=%c , c3对应的码值是%d \n", c3, c3)

	var c4 int = 22269
	fmt.Printf("c4=%c \n", c4)

	//字符类型可以进行运算，相当于一个整数，运算按照码值进行
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)
}
