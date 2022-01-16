package main

import "fmt"

//将打印金字塔代码封装到函数中

func printPyramid(totalLevel int) {

	for n := 1; n <= totalLevel; n++ {

		//打印*号前打印空格
		for k := 1; k <= totalLevel-n; k++ {
			fmt.Print(" ")
		}
		for m := 1; m <= 2*n-1; m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//乘法表
func printMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {

	var totalLevel int
	var num int

	fmt.Print("请输入金字塔行数：")
	fmt.Scanln(&totalLevel)
	printPyramid(totalLevel)

	fmt.Print("请输入乘法表层数：")
	fmt.Scanf("%d", &num)

	printMulti(num)
}
