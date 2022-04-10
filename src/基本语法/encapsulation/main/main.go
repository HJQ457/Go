package main

import (
	"fmt"
	"基本语法/encapsulation/model"
)

func main() {
	p := model.NewPerson("swith")
	p.SetAge(20)
	p.SetSalary(10000)
	fmt.Println("年龄为：", p.GetAge())
	fmt.Println("薪水为：", p.GetSalary())
}
