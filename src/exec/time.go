package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i <= 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {

	//在指定前获取时间戳
	time1 := time.Now().Unix()
	test03()

	//执行后获取时间戳
	time2 := time.Now().Unix()

	fmt.Printf("运行的时间：%v 秒", time2-time1)

}
