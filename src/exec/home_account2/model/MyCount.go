package model

import "fmt"

type familyAccount struct {
	//声明一个变量，接收用户输入
	key string
	//声明一个变量，是否退出for循环
	loop bool
	//定义账户余额
	balance float64
	//每次收支金额
	money float64
	//每次收支说明
	note string
	//收支的详情用字符串记录
	details string
	//定义一个变量，记录是否有收支行为
	flag bool
}

//编写工厂模式的构造方法，返回*FamilyAccount的实例
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

//将显示明细写成一个方法
func (this *familyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录------------------")
	if this.flag == true {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支记录，请来一笔吧")
	}
}

//登记收入写成一个方法
func (this *familyAccount) incom() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money

	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//登记支出写成一个方法
func (this *familyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)

	if this.money > this.balance {
		fmt.Println("账户余额不足")
		return
	}
	this.balance -= this.money

	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//退出系统写成一个方法
func (this *familyAccount) exit() {
	fmt.Println("确定退出吗？(y/n)")
	choice := ""

	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}

//给结构体绑定相应的方法
//显示主菜单
func (this *familyAccount) MainMenu() {
	//this.details = "收支\t账户金额\t收支金额\t说明"
	for {
		fmt.Println("\n------------------家庭收支记账软件------------------")
		fmt.Println("                  1.收支明细                      ")
		fmt.Println("                  2.登记收入                      ")
		fmt.Println("                  3.登记支出                      ")
		fmt.Println("                  4.退出软件                      ")
		fmt.Println("请选择：（1-4）")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.incom()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选项")
		}

		if !this.loop {
			break
		}
	}
}
