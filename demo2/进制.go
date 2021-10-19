package main

import "fmt"

func main() {
	var i int = 5
	//二进制输出
	fmt.Printf("i = %b\n", i)

	//八进制输出
	var j int = 011
	fmt.Println("j =", j)

	//16进制
	var k int = 0x11
	fmt.Println("k =", k)
}
