package main

import (
	"exec/fundemo03/utils"
	"fmt"
)

func main() {
	//fmt.Println(n1)	n1为局部变量，不能使用
	utils.Test()

	var n1 int = 20
	utils.Test02(n1)

	fmt.Println("(main)n1 = ", n1)

	num := 20
	utils.Test03(&num)

	a := utils.GetSum
	fmt.Printf("a的数据类型是%T\nGetSum的数据类型是%T\n", a, utils.GetSum)

	//此处的a和GetSum是等价的
	res := a(10, 40)
	fmt.Println("res = ", res)

	res2 := utils.MyFun(utils.GetSum, 50, 60)
	fmt.Println("res2 = ", res2)

	type myInt int //给int取了别名，在go中myInt和int是两个类型

	var num1 myInt
	var num2 int

	num1 = 40
	num2 = int(num1)

	fmt.Println("num = ", num1)
	fmt.Println("num2 = ", num2)

	res3 := utils.MyFun2(utils.GetSum, 500, 600)
	fmt.Println("res3 = ", res3)

	//测试可变参数的接收
	res4 := utils.Sum(10, 0, -1, 90, 10)
	fmt.Println("res4 = ", res4)

	a1 := 10
	b1 := 20
	utils.Swap(&a1, &b1)

	fmt.Printf("a1 = %v,b1 = %v\n", a1, b1)

}
