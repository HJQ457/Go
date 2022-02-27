package main

import "fmt"

func main() {
	//map声明注意事项：key不能重复
	var a map[string]string

	//在使用map时，需要先make，make的作用是给map分配数据空间
	a = make(map[string]string, 10)

	a["number1"] = "sam"
	a["number2"] = "tom"
	a["number1"] = "anna"

	fmt.Println(a)
}
