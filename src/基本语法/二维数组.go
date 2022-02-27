package main

import (
	"fmt"
)

func main() {
	var arr [4][6]int

	arr[1][2] = 1
	arr[2][1] = 1
	arr[2][3] = 1

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
