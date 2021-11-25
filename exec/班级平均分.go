package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		var sum float64

		for j := 1; j <= 3; j++ {

			var score float64

			fmt.Printf("第%v班级第%v学生成绩：\n", i, j)
			fmt.Scanf("%v", &score)

			sum += score
		}

		fmt.Println("班级平均分为：", sum/3)
	}
}
