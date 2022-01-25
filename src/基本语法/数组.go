package main

import "fmt"

func main() {
	//hen1 := 3.0
	//hen2 := 5.0
	//hen3 := 2.0
	//
	//totalWeight := hen1 + hen2 + hen3
	//avgWeight := fmt.Sprintf("%.2f",totalWeight/3)
	//
	//fmt.Printf("平均重量：%v",avgWeight)

	//定义一个数组
	var hens [3]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 2.0

	var totalWeight float64
	for i := 0; i <= len(hens)-1; i++ {
		totalWeight += hens[i]
	}

	avgWeight2 := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("平均重量为：%v", avgWeight2)
}
