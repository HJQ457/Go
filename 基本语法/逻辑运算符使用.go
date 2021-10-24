package main

import "fmt"

func main() {
	var age int = 40
	//逻辑与

	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	//逻辑或

	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//逻辑运算符使用
	if age > 30 {
		fmt.Println("ok5")
	}

	if !(age > 30) {
		fmt.Println("ok6")
	}
}
