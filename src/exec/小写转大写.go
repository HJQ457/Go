package main

import "fmt"

func main() {
	var char byte
	fmt.Println("请输入一个字符：")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	default:
		fmt.Println("other")
	}
}
