package main

import "fmt"

func main() {
	var address string = "北京长城"
	fmt.Println("address=", address)

	//字符串赋值后不可修改
	var str = "hello"
	//str[0]='a' 不可修改
	fmt.Printf("str的第一个字母是%c \n", str[0])

	//字符串的两种使用形式，用反引号原生输出，包括换行和特殊字符
	str2 := `abc\nabc`
	fmt.Println(str2)

	//字符串拼接方式
	var str1 = "hello " + "world"
	str1 += " haha!"
	fmt.Println("str1=", str1)

	//当一个拼接的操作很长时，可以分行写,但是注意“+”只能在上面的行
	var str3 = "hello " + "world " + "hello " + "world " + "hello " + "world " + "hello " + "world " +
		"hello " + "world " + "hello " + "world " + "hello " + "world " + "hello " + "world"
	fmt.Println("str3=", str3)
}
