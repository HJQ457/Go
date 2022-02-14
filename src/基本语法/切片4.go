package main

import "fmt"

func main() {
	str := "hello"
	slice := str[2:]
	fmt.Println(slice)

	//string是不可变数组，也就是说不能通过str[0] = 'z' 方式来修改字符串
	//str[0] = 'z' 是错误的
}
