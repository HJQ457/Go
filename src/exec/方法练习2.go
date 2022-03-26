package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println(i)
}

//编写一个方法，改变i的值
func (i *integer) print1() {
	*i = *i + 1
	//fmt.Println(i)
}

type sutdent struct {
	name string
	age  int
}

func (stu *sutdent) String() string {
	str := fmt.Sprintf("name=[%v],age=[%v]", stu.name, stu.age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.print1()
	fmt.Println(i)

	stu := sutdent{
		name: "tom",
		age:  10,
	}
	fmt.Println(&stu)
}
