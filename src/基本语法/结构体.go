package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	Ptr    *int
	Slice  []int
	map1   map[string]string
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.Ptr == nil {
		fmt.Println("ok1")
	}

	if p1.Slice == nil {
		fmt.Println("ok2")
	}

	//使用slice
	p1.Slice = make([]int, 2)
	p1.Slice[0] = 10
	fmt.Println(p1)

	//使用map先make
	p1.map1 = make(map[string]string)
	p1.map1["key"] = "tom"
	fmt.Println(p1)

	//不通结构体变量是独立的，互不影响，一个结构体变量字段更改不影响另一个，结构体是值类型
	var monster1 Person
	monster1.Name = "tom"
	monster1.Age = 100

	monster2 := monster1
	monster2.Name = "Sam"
	monster2.Age = 50

	fmt.Println("monster1 = ", monster1)
	fmt.Println("monster2 = ", monster2)
}
