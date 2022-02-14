package main

import "fmt"

func main() {
	var str string = "hello"
	slice := str[2:]
	fmt.Println(slice)

	//string是不可变数组，也就是说不能通过str[0] = 'z' 方式来修改字符串
	//str[0] = 'z' 是错误的
	//如果需要修改字符串，可以将string -> []byte 或者 []rune 进行修改，之后重写成string

	arr1 := []byte(str)
	arr1[1] = 'o'
	str1 := string(arr1)
	fmt.Println(str1)

	//[]byte是按照字节来处理，无法处理汉字，汉字要用[]rune处理
	arr2 := []rune(str)
	arr2[0] = '北'
	str2 := string(arr2)
	fmt.Println(str2)

}
