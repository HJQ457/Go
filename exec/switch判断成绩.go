package main

import "fmt"

func main() {
	var score float64
	fmt.Println("输入成绩：")
	fmt.Scanln(&score)

	switch int(score / 60) {
	case 1:
		fmt.Println("大于60")
	case 0:
		fmt.Println("小于60")
	default:
		fmt.Println("输入有误")
	}
}
