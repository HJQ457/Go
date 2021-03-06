package main

import "fmt"

//定义全局变量方式1
var m1 = 100
var m2 = 200
var name1 = "jack"

//定义全局变量方式2
var (
	m3    = 100
	m4    = 200
	name2 = "jack"
)

func main() {
	//golang变量使用方式1
	//1.指定变量类型，声明后不赋值，使用默认值
	//int默认值为0
	var i int
	fmt.Println("i=", i)

	//2.根据值自行判断变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	//3.省略var,注意 ：= 左侧的变量不应该是已经声明过的，否则会导致编译错误
	//下面的方式等价于 var name string name = "tom"
	name := "tom"
	fmt.Println("name", name)

	//一次声明多个变量1
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//一次声明多个变量2
	var n4, n5, n6 = 100, "tom", 888
	fmt.Println("n4=", n4, "n5=", n5, "n6=", n6)

	//一次声明多个变量3
	n7, n8, n9 := 100, "tom", 888
	fmt.Println("n7=", n7, "n8=", n8, "n9=", n9)

	//输出全局变量
	fmt.Println("m1=", m1, "m2=", m2, "name1=", name1)
	fmt.Println("m3=", m3, "m4=", m4, "name2=", name2)
}
