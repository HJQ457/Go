package main

import "fmt"

func main() {
	var num int
	var positiveCount int
	var negativeCount int

	for {
		println("请输入一个整数：")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if num > 0 {
			positiveCount++
			continue
		}

		negativeCount++
	}

	fmt.Println("整数个数为：", positiveCount)
	fmt.Println("负数个数为：", negativeCount)
}
