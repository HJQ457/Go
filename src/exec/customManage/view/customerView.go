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

//得到用户输入的id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("-------------删除客户-------------")
	fmt.Println("请选择删除的客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("-------------删除完成-------------")
		} else {
			fmt.Println("------删除失败，输入的id号不存在------")
		}
	}
}

//退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("您的输入有误，确认是否退出（Y/N）：")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
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
			this.delete()
		case "4":
			this.List()
		case "5":
			//this.loop = false
			this.exit()
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
