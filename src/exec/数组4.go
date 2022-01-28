package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [5]int

	//为了每次生成随机数，需要给一个seed值
	rand.Seed(time.Now().UnixNano())

	//定义len的长度，不在程序中反复计算，提高性能
	len := len(intArr)

	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100) //0 <= n < 100
	}

	fmt.Println("交换前：", intArr)

	fmt.Println()

	//倒叙打印方法一
	//for i := 0;i <len(intArr);i++ {
	//	fmt.Printf("intArr[%v]=%v\n",i,intArr[len(intArr) - 1 -i])
	//}

	//倒叙打印方法二
	var temp int = 0
	for i := 0; i < len/2; i++ {
		temp = intArr[i]
		intArr[i] = intArr[len-i-1]
		intArr[len-i-1] = temp
	}

	fmt.Println("交换后：", intArr)
}
