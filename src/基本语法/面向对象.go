package main

import "fmt"

//3.定义一个结构体，将cat的各个字段/属性信息，放入到Cat结构体管理
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {

	////1.使用变量处理
	//var cat1Name string = "小白"
	//var cat1Age int = 3
	//var cat1Color string = "白色"
	//
	//var cat2Name string = "小花"
	//var cat2Age int = 100
	//var cat2Color string = "花色"
	//
	////2.使用数组处理
	//var catNames [2]string = [2]string{"小白","小花"}

	//创建一个cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
}
