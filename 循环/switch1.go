package main

import "fmt"

func main() {
	var key byte
	fmt.Println("请输入：")
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("输入有误")
	}

	//switch 穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只穿透一层,不管下面是否成立，输出ok2
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("no")

	}
}
