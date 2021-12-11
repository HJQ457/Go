package main

import "fmt"

func main() {
	var n1 int32 = 10
	var n2 int32 = 50
	if n1+n2 >= 50 {
		fmt.Println("hello world")
	}

	var n3 float64 = 10.0
	var n4 float64 = 20.0
	if n3 >= 10 && n4 <= 20 {
		fmt.Println("n3 + n4 = ", n3+n4)
	}

	var num1 int32 = 10
	var num2 int32 = 5
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Println("能被3和5整除")
	}
}
