package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

func main() {
	//方式1：
	var p1 Person1
	p1.Name = "tom"
	p1.Age = 10
	fmt.Println(p1)

	//方式2：
	p2 := Person1{"Marry", 20}
	fmt.Println(p2)

	//方式3：
	var p3 *Person1 = new(Person1)
	//p3 := new(Person)
	(*p3).Name = "Sam"
	(*p3).Age = 30
	//也可以写成p3.Age = 30,go底层做了处理
	fmt.Println(*p3)

	//方式4：
	var p4 *Person1 = &Person1{}
	//(*p4).Name = "John"，go底层做了处理
	p4.Name = "John"
	p4.Age = 15
	fmt.Println(*p4)
}
