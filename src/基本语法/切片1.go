package main

import "fmt"

func main() {
	//初始化不定长数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	//声明一个切片
	//1.slice 是切片名字
	//2.intArr[1:3] 表示 slice 引用到intArr这个数组
	//3.引用 intArr 数组的起始下标为 1 ，最后的下标为3（但是不包含），左闭右开
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice=", slice)
	fmt.Println("slice len=", len(slice))

	//var slice = arr[0:end] 可以简写为 var slice = arr[:end]
	//var slice = arr[start:len(arr)] 可以简写为 var slice = arr[start:]
	//var slice = arr[0:len(arr)] 可以简写为 var slice = arr[:]
}
