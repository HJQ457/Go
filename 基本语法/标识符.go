package main

import (
	"fmt"
)

func main() {
	//Golang严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v,Num=%v", num, Num)

	//标识符中不能含有空格
	//"_"本身是go中的一个特殊标识符，称为空标识符，仅能作为占位符使用
	//不能使用系统保留关键字作为标识符
	//标识符首字母大写是公开的（public）。小写是私有（private）
}
