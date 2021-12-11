package main

import "fmt"

func main() {
	var month int32
	var age int32
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {

		if age >= 18 && age <= 60 {
			fmt.Println("60")
		} else if age < 18 {
			fmt.Println("30")
		} else {
			fmt.Println("20")
		}

	} else {

		if age >= 18 && age <= 60 {
			fmt.Println("40")
		} else {
			fmt.Println("20")
		}

	}
}
