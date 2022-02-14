package main

import "fmt"

func fbn1(n int) []int {
	//声明一个切片，切片大小为n
	fbnSlice := make([]int, n)

	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	return fbnSlice
}

func main() {
	fbnSlice := fbn1(10)
	fmt.Println(fbnSlice)
}
