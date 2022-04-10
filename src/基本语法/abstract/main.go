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
func (account *Acctract) Desposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	account.Blance = account.Blance + money
	fmt.Println("存款成功,余额为：", account.Blance)
}

//2.取款
func (account *Acctract) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 || money > account.Blance {
		fmt.Println("money error")
		return
	}
	account.Blance -= money
	fmt.Println("取款成功,余额为：", account.Blance)
}

//3.余额
func (account Acctract) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	fmt.Println("余额为：", account.Blance)
}

func main() {
	var account Acctract
	account.AcctractNo = "gs"
	account.Pwd = "123"
	account.Blance = 100

	account.Query("123")

	account.WithDraw(50, "123")

	account.Desposite(150, "123")
}
