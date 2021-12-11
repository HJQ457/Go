package main

import (
	"exec/fundemo01/utils"
	"fmt"
)

func main() {
	//输入两个数，在输入一个运算符，得到结果

	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'

	result := utils.Cal(n1, n2, operator)
	fmt.Println("result =", result)

}
