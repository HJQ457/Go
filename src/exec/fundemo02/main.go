package main

import "fmt"

func main() {
	n1 := 1
	test(n1)
	fmt.Println("(main)n1 = ", n1)

	sum := getSum(10, 20)
	fmt.Println("(main)sum = ", sum)

	res1, res2 := getSumandSub(1, 2)
	fmt.Println("res1 =", res1)
	fmt.Println("res2 = ", res2)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("(test)n1 = ", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("(getSum)sum = ", sum)
	return sum
}

//编写函数，计算和和差
func getSumandSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
