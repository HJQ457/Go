package main

import "fmt"

func main() {
	//切片使用方式1：引用数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice=", slice)
	fmt.Println("slice len=", len(slice))

	fmt.Println()

	//切片使用方式2：通过make创建切片
	//切片的长度就是它所包含的元素个数。
	//切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	var slice1 []float64 = make([]float64, 5, 10)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println("slice1=", slice1)
	fmt.Println("slice1的大小(长度):", len(slice1))
	fmt.Println("slice1的容量:", cap(slice1))

	//切片的使用方式3：定义一个切片，直接指定具体数组
	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2的长度：", len(slice2))
	fmt.Println("slice2的容量：", cap(slice2))
}
