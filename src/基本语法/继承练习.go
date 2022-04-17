package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	score float64
}

type C struct {
	A
	B
}

type D struct {
	a A //有名结构体
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type Tv struct {
	Goods
	Brand
}

func main() {
	var c C
	//c.Name = "Tom"	编译报错，因为A和B中都有Name
	c.A.Name = "Tom" //必须通过匿名结构体来区分
	fmt.Println(c.A.Name)

	var d D
	//d.name = "jack"	有名结构体不会向上面找
	d.a.Name = "jack"
	fmt.Println(d.a.Name)

	tv := Tv{Goods{"电视机01", 5000.99}, Brand{"海尔", "山东青岛"}}
	tv2 := Tv{
		Goods{
			Price: 5000.96,
			Name:  "电视机02",
		},
		Brand{
			Name:    "海尔",
			Address: "北京",
		},
	}
	fmt.Println(tv)
	fmt.Println(tv2)
}
