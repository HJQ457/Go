package main

import "fmt"

func main() {
	var age int

	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)

	if age > 18 {
		fmt.Println("年龄大于18")
	} else {
		fmt.Println("年龄小于18")
	}

	if age := 20; age > 18 { //golang支持在if中，直接定义一个变量
		fmt.Println("年龄大于18")
	} else {
		fmt.Println("年龄小于18")
	}
}
