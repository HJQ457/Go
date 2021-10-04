package main

import "fmt"

func main() {
	var price float32 = 89.12

	//尾部部分可能出现精度损失,双精度
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901

	//单精度float32  4字节  -3.403E38~3.403E38
	//双精度float64  8字节  -1.798E308~1.798E308

	fmt.Println("price=", price)

	fmt.Println("num1=", num1, "num2=", num2)

	//浮点型默认为float64
	var num5 = 1.1
	fmt.Printf("num5数据类型是 %T \n", num5)

	//十进制形式
	num6 := 5.12
	num7 := .123 // =0.123

	fmt.Println("num6=", num6, "num7=", num7)

	//科学计数法
	num8 := 5.1234e2 // = 5.1234 * 10^2

	fmt.Println("num8=", num8)
}
