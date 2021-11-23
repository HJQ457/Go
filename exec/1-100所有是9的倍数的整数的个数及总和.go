package main

import "fmt"

func main() {
	var sum int = 0
	var count int = 0

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum = sum + i
			count = count + 1
		}
	}

	fmt.Println("数量为：", count)
	fmt.Println("总和为：", sum)
}
