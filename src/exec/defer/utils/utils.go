package utils

import "fmt"

func Sum(n1 int, n2 int) int {

	//当执行到defer时，暂时不执行，会将后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 res = ", res)
	return res
}
