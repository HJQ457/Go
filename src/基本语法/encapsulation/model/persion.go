package model

import "fmt"

type person struct {
	Name   string
	age    int //其他包不能访问
	salary float64
}

//工厂模式函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和salary，需要写一个方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入错误")
	}
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary < 30000 {
		p.salary = salary
	} else {
		fmt.Println("输入错误")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

func (p *person) GetAge() int {
	return p.age
}
