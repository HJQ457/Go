package main

import "fmt"

func main() {
	var name string
	var passwd string
	var logChance byte = 3

	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanf("%v", &name)
		fmt.Println("请输入密码：")
		fmt.Scanf("%v", &passwd)

		if name == "小明" && passwd == "888" {
			println("登入成功")
			break
		} else {
			logChance--
			fmt.Printf("您还有%v次机会\n", logChance)

		}

		if logChance == 0 {
			fmt.Println("登录失败")
		}
	}
}
