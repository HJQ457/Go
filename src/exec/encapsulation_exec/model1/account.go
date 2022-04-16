package model1

import "fmt"

type account struct {
	accountNum string
	pwd        string
	balance    float64
}

//工厂模式，将私有化暴露
func Newaccount(accountNum string, pwd string, balance float64) *account {

	if len(accountNum) < 6 || len(accountNum) > 20 {
		fmt.Println("账号长度不对")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度不正确")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额不正确")
		return nil
	}

	return &account{
		accountNum: accountNum,
		pwd:        pwd,
		balance:    balance,
	}
}

//方法
//1.查余额
func (acc account) Find(num string, pwd string) {
	if acc.pwd != pwd || acc.accountNum != num {
		fmt.Println("密码或密码错误")
	} else {
		fmt.Println("余额为:", acc.balance)
	}
}

//1.取款
func (acc account) Qukuan(num string, pwd string, money float64) {
	if acc.pwd != pwd || acc.accountNum != num {
		fmt.Println("密码或密码错误")
		return
	}

	if acc.balance < money || money < 0 {
		fmt.Println("余额不足或取款金额不正确")
		return
	}

	acc.balance -= money
	fmt.Println("取款成功，余额为：", acc.balance)
}

//2.存款
func (acc account) Cunkuan(num string, pwd string, money float64) {
	if acc.pwd != pwd || acc.accountNum != num {
		fmt.Println("密码或密码错误")
		return
	}

	if money < 0 {
		fmt.Println("金额不正确")
		return
	}

	acc.balance += money
	fmt.Println("取款成功，余额为：", acc.balance)
}
