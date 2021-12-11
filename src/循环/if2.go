package main

import "fmt"

func main() {
	var score float32
	fmt.Println("请输入分数：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("iphone")
	} else if score >= 60 && score <= 80 {
		fmt.Println("ipad")
	} else {
		fmt.Println("nothing")
	}
}
