package main

import "fmt"

func main() {
	//接收用户输入
	var key int

	//判断是否循环显示菜单
	var loop = true

	var userId int
	var userPwd string

	for loop {
		fmt.Println("----------欢迎登入多人聊天系统----------")
		fmt.Println("\t 1.登录聊天室")
		fmt.Println("\t 2.注册用户")
		fmt.Println("\t 3.退出系统")
		fmt.Println("\t 请选择（1-3）：")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

	//提示信息
	if key == 1 {
		fmt.Println("请输入用户id：")
		fmt.Scanln(&userId)
		fmt.Println("请输入密码：")
		fmt.Scanln(&userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
			return
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册")
	}
}
