package main

import (
	"fmt"
)

func main() {
	names := []string{"sam", "tom", "jack", "anna"}

	fmt.Println("请输入要查找的人名：")
	var name string
	fmt.Scanf("%v", &name)

	index := -1

	for i := 0; i < len(names); i++ {
		if names[i] == name {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("下标为%v", index+1)
	} else {
		fmt.Println("没有找到")
	}
}
