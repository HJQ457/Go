package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i=", i)

	//int8   1字节  -128~127
	//int16  2字节  -2^15~2^15-1
	//int32  4字节  -2^31~2^31-1
	//int64  6字节  -2^63~2^63-1

	//uint8  1字节  0~255
	//uint16 2字节  0~2^16-1
	//uint32 4字节  0~2^32-1
	//uint64 8字节  0~2^64-1

	//int	 32位系统4字节  -2^31~2^31-1
	//int	 64位系统4字节  -2^63~2^63-1

	//uint	 32位系统4字节  0~2^32-1
	//uint	 32位系统8字节  0~2^64-1

	//rune	 4字节		  -2^31~2^31-1

	//byte   1字节		  0~255

	//测试int8范围
	var j int8 = -128
	fmt.Println("j=", j)

	//测试uint8范围
	var k uint8 = 0
	fmt.Println("k=", k)

	//测试byte范围
	var c byte = 255
	fmt.Println("c=", c)

	//整型的使用细节，格式化输出
	var n1 int = 100
	fmt.Printf("n1 的类型 %T \n", n1)

	//查看变量字节占用大小和数据类型
	var n2 int64 = 10
	//unsafe.Sizeof() 是unsafe包的一个函数，可以返回变量占用的字节数
	fmt.Printf("n2 的数据类型 %T，n2的字节占用数是 %d", n2, unsafe.Sizeof(n2))
}
