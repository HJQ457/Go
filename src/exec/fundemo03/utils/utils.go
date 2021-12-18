package utils

import "fmt"

func Test() {
	var n1 int = 10
	fmt.Println("n1 =", n1)
}

func Test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("(test02)n1 = ", n1)
}

//n1 是*int指针
func Test03(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("(test03)n1 = ", *n1)
}

func GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数是一种基本数据类型，因此在GO中，可以作为形参，并且调用
func MyFun(funcvar func(int, int) int, num1 int, num2 int) int {
	return funcvar(num1, num2)
}

type MyFunType func(int, int) int

func MyFun2(funcvar MyFunType, num1 int, num2 int) int {
	return funcvar(num1, num2)
}

//编写函数，求多个int的和
func Sum(n1 int, args ...int) int {
	sum := n1

	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

//编写一个函数交换n1和n2的值
func Swap(n1 *int, n2 *int) {
	//t := *n1
	//*n1 = *n2
	//*n2 = t

	*n1, *n2 = *n2, *n1
}
