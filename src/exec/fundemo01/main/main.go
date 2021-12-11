package main

import (
	//"exec/fundemo01/utils"
	util "exec/fundemo01/utils" //别名
	"fmt"
)

func main() {
	//输入两个数，在输入一个运算符，得到结果

	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'

	//result := utils.Cal(n1, n2, operator)
	result := util.Cal(n1, n2, operator)
	fmt.Println("result =", result)

	//fmt.Println("utils.go 的 num:",utils.Num1)
	fmt.Println("utils.go 的 num:", util.Num1)

}
