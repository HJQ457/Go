package main

import "fmt"

func main() {
	var score1 [5]float64

	for i := 0; i < len(score1); i++ {
		fmt.Printf("请输入第%v个数字\n", i+1)
		fmt.Scanf("%v\n", &score1[i])
	}

	for i := 0; i < len(score1); i++ {
		fmt.Printf("score=%v\n", score1[i])
	}
}
