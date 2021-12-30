package main

import "fmt"

//作用域在整个包都有效，如果首字母大写，则作用域在整个程序有效
var age1 int = 50
var Name string = "jack~"

func test3() {
	//age和Name的作用域就只在test函数内部
	age := 10
	Name := "tom~"
	fmt.Println("age = ", age)
	fmt.Println("Name = ", Name)
}
func main() {
	fmt.Println("age = ", age1)
	fmt.Println("Name = ", Name)
	fmt.Println()
	test3()
}
