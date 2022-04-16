package main

import "exec/encapsulation_exec/model1"

func main() {
	acc := model1.Newaccount("123456", "123123", 50)

	acc.Find("123456", "123123")
	acc.Qukuan("123456", "123123", 20)
	acc.Find("123456", "123123")
	acc.Cunkuan("123456", "123123", 0)

}
