package main

import "fmt"

func main() {
	//基本数据类型在内存布局
	var i int = 10
	//i的地址，&i
	fmt.Println("i的地址=", &i)

	//var ptr *int = &i
	//1.ptr是指针变量
	//2.ptr的类型 *int
	//3.ptr本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr=%v\n", &ptr)
	fmt.Printf("ptr=%v\n", *ptr)

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr1 *int = &num
	fmt.Printf("num=%v\n", num)
	//ptr1 = &num
	*ptr1 = 10
	fmt.Printf("num=%v\n", num)
}
