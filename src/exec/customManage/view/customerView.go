package main

import (
	"exec/customManage/model"
	"exec/customManage/service"
	"fmt"
)

type customerView struct {
	key             string //接收用户输入
	loop            bool   //是否循环显示主菜单
	customerService *service.CustomerService
}

//显示所有客户信息
func (this *customerView) List() {

	//获取当前所有的客户信息
	customers := this.customerService.List()

	//显示
	fmt.Println("-------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].Getinfo())
	}
	fmt.Println("------------客户列表完成------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("-------------添加客户-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	customers := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customers) {
		fmt.Println("-------------添加完成-------------")
	} else {
		fmt.Println("-------------添加失败-------------")
	}
}

//显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-------------客户信息管理软件-------------")
		fmt.Println("             1.添加客户                 ")
		fmt.Println("             2.修改客户                 ")
		fmt.Println("             3.删除客户                 ")
		fmt.Println("             4.客户列表                 ")
		fmt.Println("             5.退   出                 ")
		fmt.Print("请选择：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("2")
		case "3":
			fmt.Println("3")
		case "4":
			this.List()
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了软件")
}
func main() {
	customcerView := customerView{
		key:  "",
		loop: true,
	}

	//这里完成对customerView结构体的customerService字段的初始化
	customcerView.customerService = service.NewCustomerService()

	customcerView.mainMenu()
}
