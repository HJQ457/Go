package main

import (
	"fmt"
)

func main() {
	var arr1 = [][]int{{1, 2, 3}, {4, 5, 6}}

	//for循环遍历
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%v\t", arr1[i][j])
		}
		fmt.Println()
	}

	//for-range遍历
	for i, v := range arr1 {
		for j, v2 := range v {
			fmt.Printf("arr1[%v][%v]=%v \n", i, j, v2)
		}
	}
}
