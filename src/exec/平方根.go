package main

import (
	"fmt"
	"math"
)

//解方程 ax^2 + bx + c = 0
func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("输入a：")
	fmt.Scanln(&a)

	fmt.Println("输入b：")
	fmt.Scanln(&b)

	fmt.Println("输入c：")
	fmt.Scanln(&c)

	if b*b-4*a*c > 0 {
		var x1 float64 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
		var x2 float64 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

		fmt.Printf("x1 = %v,x2 = %v", x1, x2)
	} else if b*b-4*a*c == 0 {
		var x1 float64 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
		fmt.Println("x1 = x2 =", x1)
	} else {
		fmt.Println("无解")
	}
}
