package main

import (
	"fmt"
	"strconv"
)

func main() {
	//统计字符串长度
	str := "hello北"
	//golang字符串编码统一为utf-8,len函数获取的是字节，不是字符
	fmt.Println(len(str))

	str2 := "hello北京"

	//字符串遍历，处理中午问题 r := []rune(str)
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符 = %c\n", str3[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果：", n)
	}

	//整数转字符串
	str4 := strconv.Itoa(12345)
	fmt.Printf("str4 = %v,str = %T", str4, str4)
}
