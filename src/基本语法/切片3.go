package main

import "fmt"

func main() {
	var slice3 []int = []int{100, 200, 300}

	//用append内置函数，可以对切片进行动态增加
	slice3 = append(slice3, 400, 500, 600)

	slice3 = append(slice3, slice3...)

	fmt.Println("slice3=", slice3)

	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)
}
