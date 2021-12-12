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
