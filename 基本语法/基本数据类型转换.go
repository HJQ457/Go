package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int32 = 100
	//将i转换成float
	var j float32 = float32(i)
	var k int8 = int8(i)

	fmt.Printf("i=%v j=%v k=%v \n", i, j, k)
	fmt.Printf("i的数据类型是	%T\n", i)

	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1 + 20)
	n3 = int8(n1 + 20)

	fmt.Println("n2=", n2, "n3=", n3)

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	//使用第一种方式转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T,str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T,str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T,str=%v\n", b, b)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T,str=%q\n", mychar, mychar)

	//第二种方式strconv函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10) //base为10 表示转换成10进制
	fmt.Printf("str type %T,str=%v\n", str, str)

	//f代表格式 10：小数位保留10位 64：表示这个
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T,str=%v\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T,str=%v\n", str, str)

	//strconv包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T,str=%v\n", str, str)
}
