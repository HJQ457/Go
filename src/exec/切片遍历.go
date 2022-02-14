package main

import "fmt"

func main() {
	var arr1 [5]int = [...]int{10, 20, 30, 40, 50}
	slice1 := arr1[1:4]

	//使用for循环遍历切片
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%v]=%v\n", i, slice1[i])
	}

	fmt.Println()

	//使用for--range遍历切片
	for i, v := range slice1 {
		fmt.Printf("slice1[%v]=%v\n", i, v)
	}
}
