package main

import "fmt"

//结构体
type Acctract struct {
	AcctractNo string
	Pwd        string
	Blance     float64
}

//方法
//1.存款
func (account Acctract) Desposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	account.Blance = account.Blance + money
}

func main() {

}
