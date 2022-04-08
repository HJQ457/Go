package main

import (
	"fmt"
)

type MethodUtils1 struct {
	num1 float64
	num2 float64
}

func (m *MethodUtils1) Prepare(num int) {
	if num%2 == 0 {
		fmt.Printf("%v是偶数", num)
	} else {
		fmt.Printf("%v是奇数", num)
	}
}

func (m1 *MethodUtils1) Print(m int, n int, key string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func (m2 *MethodUtils1) Sub(m float64, n float64) float64 {
	return m + n
}

func (m3 *MethodUtils1) Oper(oper byte) {
	var res float64
	switch oper {
	case '+':
		res = m3.num1 + m3.num2
	case '-':
		res = m3.num1 - m3.num2
	case '*':
		res = m3.num1 * m3.num2
	case '/':
		res = m3.num1 / m3.num2
	}
	fmt.Println("res=", res)
}

func main() {
	var m MethodUtils1
	m.Prepare(3)
	fmt.Println()
	(&m).Prepare(2) //严格来说，应该这么写
	fmt.Println()

	var m1 MethodUtils1
	m1.Print(3, 5, "#")

	var m3 MethodUtils1
	fmt.Printf("sum=%v", fmt.Sprintf("%.2f", m3.Sub(1.2, 2.5))) //处理精度问题
	fmt.Println()
	fmt.Println(fmt.Sprintf("%.5f", m3.Sub(2.1, 3.6)))

	var m4 MethodUtils1
	m4.num1 = 2
	m4.num2 = 3
	m4.Oper('*')
}
