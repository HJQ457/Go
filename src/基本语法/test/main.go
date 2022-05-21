package main

import "fmt"

//一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	//传统测试方法
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper err,返回值%v 期望值%v", res, 55)
	} else {
		fmt.Println("addUpper true,返回值%v 期望值%v", res, 55)
	}
}
