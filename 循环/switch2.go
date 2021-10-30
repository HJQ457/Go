package main

import "fmt"

func main() {
	var n1 int32 = 20
	var n2 int32 = 20

	//switch 和 case 的数据类型要保持一致

	switch n1 {
	case n2, 10, 5: //case 后面可以有多个表达式
		fmt.Println("ok")

	//case 5:	前面有常量5，会报错

	default:
		fmt.Println("没有匹配到")
	}

	var score int = 30
	switch {
	case score > 90:
		fmt.Println("A")
	case score >= 70 && score <= 90:
		fmt.Println("B")
	case score >= 60 && score < 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

}
