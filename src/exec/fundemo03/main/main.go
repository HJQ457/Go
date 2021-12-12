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
}
