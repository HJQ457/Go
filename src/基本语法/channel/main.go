package main

import "fmt"

func main() {
	var intChar chan int        //chan是关键字
	intChar = make(chan int, 3) //管道必须make才能使用

	//输出intChar
	fmt.Printf("intChar值为%v，intChar本身的地址%p\n", intChar, &intChar)

	//写入数据
	intChar <- 10
	var num int = 211
	intChar <- num

	//查看容量
	fmt.Printf("channel len = %v，cap = %v\n", len(intChar), cap(intChar))

	//从管道中读取数据
	var num2 int
	num2 = <-intChar

	//管道是先进先出，所以是10
	fmt.Println(num2)

	num3 := <-intChar
	fmt.Println(num3)
}
