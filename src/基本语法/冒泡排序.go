package main

import "fmt"

func BubbleSort(arr *[]int) {
	fmt.Println("排序前数组：", *arr)

	//完成排序
	for i := 0; i < len(*arr)-1; i++ {
		//完成第一轮排序
		var temp int
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后数组：", *arr)
}

func main() {
	//定义数组
	//arr := [6]int{24,69,80,57,13,15}

	//定义切片
	arr := []int{24, 69, 80, 57, 13, 15}
	//将数组传递给一个函数，完成排序
	BubbleSort(&arr)
}
