package model

import "fmt"

//声明一个customer结构体，表示客户信息
type Customer struct {
	Id     int
	Name   string
	Gerder string
	Age    int
	Phone  string
	Email  string
}

//编写工厂模式，返回customer实例
func NewCustomer(id int, name string, gerder string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gerder: gerder,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建customer实例的方法，不带id
func NewCustomer2(name string, gerder string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gerder: gerder,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户信息
func (this Customer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gerder, this.Age, this.Phone, this.Email)
	return info
}
