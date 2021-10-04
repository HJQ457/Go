package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b)

	//布尔类型占用一个字节
	fmt.Println("b的占用空间", unsafe.Sizeof(b))

	//bool类型只能取true或者false
}
