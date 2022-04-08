package main

import (
	"fmt"
)

type Box struct {
	long   float64
	weight float64
	heignt float64
}

func (b *Box) Tiji() float64 {
	tiji := (*b).long * (*b).weight * (*b).heignt
	return tiji
}

type Visitor struct {
	name string
	age  int
}

func (v *Visitor) showPrice() {
	if (*v).age > 18 {
		fmt.Printf("游客名字为：%v，年龄为：%v，收费为：20", (*v).name, (*v).age)
	} else {
		fmt.Printf("游客名字为：%v，年龄为：%v，收费为：10", (*v).name, (*v).age)
	}
}

func main() {
	var b Box

	fmt.Println("长：")
	fmt.Scanln(&b.long)
	fmt.Println("宽：")
	fmt.Scanln(&b.weight)
	fmt.Println("高：")
	fmt.Scanln(&b.heignt)

	fmt.Println("体积为：", (&b).Tiji())

	var v Visitor
	v.name = "jack"
	v.age = 19
	(&v).showPrice()
}
