package main

import "fmt"

var age = test2()

func test2() int {
	fmt.Println("test()...")
	return 90
}

//通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...", age)
}
